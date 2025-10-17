from PIL import Image, ImageDraw, ImageFont
from io import BytesIO
import cairosvg
from pathlib import Path
from .guilloche import generate_guilloche_png, generate_sunburst_png, generate_rosette_png, generate_concentric_circles_png, generate_wavy_rings_png

# Pick a font and size
FONT_PATH = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"
FONT_MONO_PATH = "/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf"
FONT_CJK_PATH = "/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc"
large_font = ImageFont.truetype(FONT_PATH, 56)     # xlarge
try:
    title_font = ImageFont.truetype(FONT_CJK_PATH, 32)
except Exception:
    title_font = ImageFont.truetype(FONT_PATH, 32)
label_font = ImageFont.truetype(FONT_PATH, 20)     # medium
mono_large_font = ImageFont.truetype(FONT_MONO_PATH, 56)
mono_label_font = ImageFont.truetype(FONT_MONO_PATH, 20)


IMAGE_SIZE = (580, 1000)
IMAGE_BACKGROUND_COLOR = (170, 241, 231, 100)

def make_linear_gradient(size, start_rgba, end_rgba, horizontal=False) -> Image.Image:
    width, height = size
    if horizontal:
        grad = Image.new("RGBA", (width, 1))
        for x in range(width):
            t = x / (width - 1) if width > 1 else 0
            r = int(start_rgba[0] + (end_rgba[0] - start_rgba[0]) * t)
            g = int(start_rgba[1] + (end_rgba[1] - start_rgba[1]) * t)
            b = int(start_rgba[2] + (end_rgba[2] - start_rgba[2]) * t)
            a = int(start_rgba[3] + (end_rgba[3] - start_rgba[3]) * t)
            grad.putpixel((x, 0), (r, g, b, a))
        return grad.resize(size)
    else:
        grad = Image.new("RGBA", (1, height))
        for y in range(height):
            t = y / (height - 1) if height > 1 else 0
            r = int(start_rgba[0] + (end_rgba[0] - start_rgba[0]) * t)
            g = int(start_rgba[1] + (end_rgba[1] - start_rgba[1]) * t)
            b = int(start_rgba[2] + (end_rgba[2] - start_rgba[2]) * t)
            a = int(start_rgba[3] + (end_rgba[3] - start_rgba[3]) * t)
            grad.putpixel((0, y), (r, g, b, a))
        return grad.resize(size)


def wrap_text(text: str, font: ImageFont.FreeTypeFont, max_width: int) -> list[str]:
    # Greedy word-wrapping with fallback to character-splitting for long tokens
    tmp_img = Image.new("RGBA", (1, 1))
    d = ImageDraw.Draw(tmp_img)

    words = text.split(" ")
    if not words:
        return []

    lines: list[str] = []
    current = ""
    for word in words:
        candidate = word if not current else current + " " + word
        w = d.textlength(candidate, font=font)
        if w <= max_width:
            current = candidate
        else:
            if current:
                lines.append(current)
                current = word
            else:
                # word itself too long, split by characters
                token = word
                buf = ""
                for ch in token:
                    cw = d.textlength(buf + ch, font=font)
                    if cw <= max_width:
                        buf += ch
                    else:
                        if buf:
                            lines.append(buf)
                        buf = ch
                if buf:
                    current = buf
                else:
                    current = ""
    if current:
        lines.append(current)
    return lines


def draw_rotated_text(base_img, text, xy, angle_deg, font, fill, pad=12):
    # measure tight bbox
    tmp = Image.new("RGBA", (1, 1))
    d = ImageDraw.Draw(tmp)
    x0, y0, x1, y1 = d.textbbox((0, 0), text, font=font, anchor=None, stroke_width=0)
    w, h = (x1 - x0), (y1 - y0)

    # render on padded canvas
    txt = Image.new("RGBA", (w + 2 * pad, h + 2 * pad), (0, 0, 0, 0))
    td = ImageDraw.Draw(txt)
    td.text((pad - x0, pad - y0), text, font=font, fill=fill)

    # rotate with expand, then paste with alpha
    txt_rot = txt.rotate(angle_deg, expand=True, resample=Image.BICUBIC)
    base_img.paste(txt_rot, xy, txt_rot)


def gen_image(
    book_nft_address: str,
    staked_amount: str,
    reward_index: str,
    initial_staker: str,
    book_nft_name: str | None = None,
) -> Image.Image:
    """
    Generate an image with the given parameters.

    Args:
        `book_nft_address`: The address of the book NFT.
            It is a 42 characters long string e.g. `0x1234567890123456789012345678901234567890`
        `staked_amount`: The amount of staked tokens.
            It is a string representing a number e.g. `123`
        `reward_index`: The reward index.
            It is a string representing a number e.g. `123`
        `initial_staker`: The address of the initial staker.
            It is a 42 characters long string e.g. `0x1234567890123456789012345678901234567890`
        `book_nft_name`: The name of the book NFT.

    Returns:
        An image with the given parameters.
    """
    background = make_linear_gradient(IMAGE_SIZE, (210, 240, 240, 255), (240, 230, 180, 255), horizontal=False)
    image = background

    draw = ImageDraw.Draw(image)

    if book_nft_name:
        max_title_width = IMAGE_SIZE[0] - 80 - 80  # right margin 80
        lines = wrap_text(book_nft_name, title_font, max_title_width)
        # Measure line height
        tmp = Image.new("RGBA", (1, 1))
        td = ImageDraw.Draw(tmp)
        _, y0, _, y1 = td.textbbox((0, 0), "Ay", font=title_font)
        line_height = (y1 - y0) + 4
        y = 420
        for line in lines:
            draw.text((80, y), line, fill=(40, 100, 110, 255), font=title_font)
            y += line_height + 4
    draw_rotated_text(
        image,
        f"Book - {book_nft_address}",
        xy=(535, 40),
        angle_deg=-90,
        font=label_font,
        fill=(40, 100, 110, 255),
        pad=16,
    )
    draw_rotated_text(
        image,
        f"Staker - {initial_staker}",
        xy=(0, 300),
        angle_deg=90,
        font=mono_label_font,
        fill=(40, 100, 110, 255),
        pad=16,
    )
    draw.text((80, IMAGE_SIZE[1] - 240), f"{staked_amount}", fill=(40, 100, 110, 255), font=mono_large_font)
    draw.text((80, IMAGE_SIZE[1] - 176), f"LIKE Staked", fill=(40, 100, 110, 255), font=label_font)
    draw.text((80, IMAGE_SIZE[1] - 128), f"Reward Index: {reward_index}", fill=(40, 100, 110, 255), font=label_font)

    # rectangle coordinates (x0, y0, x1, y1)
    bbox = (40, 40, 540, 960)
    draw.rounded_rectangle(bbox, radius=16, outline=(40, 100, 110, 255), width=2)

    # Guilloché overlay seeded by book_nft_address
    inner_width = 1800
    inner_height = 1800
    # Sunburst background + multi-ring lines on top
    sunburst_png = generate_sunburst_png(
        book_nft_address.lower(),
        inner_width,
        inner_height,
        rays=720,
        inner_ratio=0.002,
        outer_ratio=0.24,
        stroke_width=0.8,
        opacity=0.17,
    )
    sunburst_img = Image.open(BytesIO(sunburst_png)).convert("RGBA")
    image.alpha_composite(sunburst_img, dest=(-200, -320))

    rings_png = generate_wavy_rings_png(
        book_nft_address.lower(),
        inner_width,
        inner_height,
        stroke_width=1.0,
        opacity=0.17,
        rings=40,
        ring_spacing=30,
    )
    rings_img = Image.open(BytesIO(rings_png)).convert("RGBA")
    image.alpha_composite(rings_img, dest=(-200, -320))

    # Render and paste LikeCoin logo SVG
    svg_path = Path(__file__).resolve().parents[1] / "images" / "likecoin-logo.svg"
    if svg_path.exists():
        # Adjust logo size as needed via output_width/output_height
        logo_png_bytes = cairosvg.svg2png(url=str(svg_path), output_width=128)
        logo_img = Image.open(BytesIO(logo_png_bytes)).convert("RGBA")
        # Paste with alpha channel as mask
        image.paste(logo_img, (70 , 240), logo_img)

    return image
