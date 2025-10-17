import hashlib
import math
from typing import Tuple

import cairosvg


def _hash_to_ints(seed: str) -> Tuple[int, int, int, float]:
    h = hashlib.sha256(seed.encode("utf-8")).digest()
    a = h[0] % 9 + 3  # 3..11
    b = h[1] % 9 + 3  # 3..11
    k = h[2] % 7 + 3  # 3..9
    phase = (h[3] / 255.0) * math.tau
    return a, b, k, phase


def _generate_path(cx: float, cy: float, ax: float, ay: float, a: int, b: int, k: int, phase: float, turns: int = 8, steps: int = 2400) -> str:
    # Lissajous-hypotrochoid hybrid for guilloché-like curve
    pts = []
    total = math.tau * turns
    for i in range(steps + 1):
        t = total * (i / steps)
        r = 1.0 + 0.15 * math.sin(k * t)
        x = cx + r * ax * math.sin(a * t + phase)
        y = cy + r * ay * math.sin(b * t)
        pts.append((x, y))

    if not pts:
        return ""

    d = [f"M {pts[0][0]:.3f},{pts[0][1]:.3f}"]
    for (x, y) in pts[1:]:
        d.append(f"L {x:.3f},{y:.3f}")
    d.append("Z")
    return " ".join(d)


def generate_guilloche_png(
    seed: str,
    width: int,
    height: int,
    stroke_color: str = "#28646E",
    stroke_width: float = 2,
    opacity: float = 0.18,
    rings: int = 3,
    ring_spacing: int = 2,
) -> bytes:
    cx, cy = width / 2.0, height / 2.0

    # Build multiple rings with slight parameter variations derived from seed
    g_parts = [
        f"<svg width=\"{width}\" height=\"{height}\" viewBox=\"0 0 {width} {height}\" xmlns=\"http://www.w3.org/2000/svg\">",
        f"<g fill=\"none\" stroke=\"{stroke_color}\" stroke-linejoin=\"round\" stroke-linecap=\"round\">",
    ]

    base_a, base_b, base_k, base_phase = _hash_to_ints(seed)

    max_margin = 8 + ring_spacing * max(0, rings - 1)
    for i in range(rings):
        margin = 8 + i * ring_spacing
        ax = (width / 2.0) - margin
        ay = (height / 2.0) - margin

        # Slightly vary parameters per ring, cycling through hash
        a = base_a + (i % 3) - 1
        b = base_b + ((i * 2) % 3) - 1
        k = max(2, base_k + (i % 2))
        phase = (base_phase + (i * math.pi / 7.0)) % math.tau

        path_d = _generate_path(cx, cy, ax, ay, a, b, k, phase)

        # Taper stroke and opacity the further from center
        ring_opacity = max(0.06, opacity * (1.0 - (margin / (width / 2.0))))
        ring_width = max(0.4, stroke_width * (1.0 - (margin / (max(width, height) / 2.0))))

        g_parts.append(
            f"<path d=\"{path_d}\" stroke-width=\"{ring_width:.3f}\" opacity=\"{ring_opacity:.3f}\"/>"
        )

    g_parts.append("</g>")
    g_parts.append("</svg>")
    svg = "\n".join(g_parts)
    return cairosvg.svg2png(bytestring=svg.encode("utf-8"))



def generate_sunburst_png(
    seed: str,
    width: int,
    height: int,
    stroke_color: str = "#28646E",
    stroke_width: float = 1.2,
    opacity: float = 0.22,
    rays: int = 160,
    inner_ratio: float = 0.12,
    outer_ratio: float = 0.48,
    wobble_freq: int = 9,
    wobble_amp: float = 0.18,
) -> bytes:
    """Generate a sunburst made of radial rays seeded by the input string.

    Rays length is modulated by a sinusoid and a seed-derived jitter for variety.
    """
    cx, cy = width / 2.0, height / 2.0
    base = min(width, height)
    r_inner = base * inner_ratio
    r_outer = base * outer_ratio

    h = hashlib.sha256(seed.encode("utf-8")).digest()

    # derive simple per-ray jitter from the hash bytes
    jitters = [(h[i % len(h)] / 255.0 - 0.5) * 2.0 for i in range(rays)]  # [-1,1]

    parts = [
        f"<svg width=\"{width}\" height=\"{height}\" viewBox=\"0 0 {width} {height}\" xmlns=\"http://www.w3.org/2000/svg\">",
        f"<g fill=\"none\" stroke=\"{stroke_color}\" stroke-linecap=\"round\" opacity=\"{opacity}\">",
    ]

    for i in range(rays):
        t = (i / rays) * math.tau
        # Modulate ray length with sinusoid and jitter
        mod = 0.5 + 0.5 * math.sin(wobble_freq * t + (h[0] / 255.0) * math.tau)
        mod = (1.0 - wobble_amp) + wobble_amp * mod
        r_end = r_inner + (r_outer - r_inner) * mod * (1.0 + 0.08 * jitters[i])

        x1 = cx + r_inner * math.cos(t)
        y1 = cy + r_inner * math.sin(t)
        x2 = cx + r_end * math.cos(t)
        y2 = cy + r_end * math.sin(t)

        # Slight stroke tapering across rays
        sw = max(0.6, stroke_width * (0.8 + 0.2 * jitters[i]))
        parts.append(
            f"<line x1=\"{x1:.2f}\" y1=\"{y1:.2f}\" x2=\"{x2:.2f}\" y2=\"{y2:.2f}\" stroke-width=\"{sw:.2f}\"/>"
        )

    parts.append("</g>")
    parts.append("</svg>")
    svg = "\n".join(parts)
    return cairosvg.svg2png(bytestring=svg.encode("utf-8"))


def generate_rosette_png(
    seed: str,
    width: int,
    height: int,
    stroke_color: str = "#28646E",
    stroke_width: float | None = None,
    opacity: float | None = None,
    petals: int | None = None,
    inner_scale: float | None = None,
    outer_scale: float | None = None,
    steps: int = 3200,
    rings: int = 1,
) -> bytes:
    """Elegant rosette using a hypotrochoid-like curve with smooth modulation.

    x(t) = cx + (R - r) * cos(t) + d * cos(((R - r)/r) * t)
    y(t) = cy + (R - r) * sin(t) - d * sin(((R - r)/r) * t)

    Parameters R, r, d are derived from the seed and bounding box.
    """
    cx, cy = width / 2.0, height / 2.0
    base = min(width, height)

    h = hashlib.sha256(seed.encode("utf-8")).digest()

    # Randomize parameters from seed when not provided
    if petals is None:
        petals = 6 + (h[0] % 19)  # 6..24
    if inner_scale is None:
        inner_scale = 0.06 + (h[1] / 255.0) * 0.12  # 0.06..0.18
    if outer_scale is None:
        outer_scale = 0.40 + (h[2] / 255.0) * 0.10  # 0.40..0.50
    if stroke_width is None:
        stroke_width = 0.8 + (h[3] / 255.0) * 0.8  # 0.8..1.6
    if opacity is None:
        opacity = 0.16 + (h[4] / 255.0) * 0.14  # 0.16..0.30

    # number of concentric rosette rings (seeded)
    rings = max(1, rings)
    seeded_extra_rings = 1 + (h[5] % 2)  # 1..2 extra
    rings_total = min(1 + seeded_extra_rings, max(1, rings)) if rings > 1 else rings

    parts = [
        f"<svg width=\"{width}\" height=\"{height}\" viewBox=\"0 0 {width} {height}\" xmlns=\"http://www.w3.org/2000/svg\">",
        f"<g fill=\"none\" stroke=\"{stroke_color}\" stroke-linejoin=\"round\" stroke-linecap=\"round\">",
    ]

    for i in range(rings_total):
        # Per-ring slight variations
        ring_outer = outer_scale * (1.0 - 0.04 * i)
        ring_inner = max(0.02, inner_scale * (1.0 - 0.06 * i))
        ring_petals = max(5, petals + ((h[6 + i] % 5) - 2))
        ring_sw = max(0.5, stroke_width * (1.0 - 0.1 * i))
        ring_opacity = max(0.08, opacity * (1.0 - 0.12 * i))
        ring_rot = (h[10 + i] / 255.0) * math.tau  # rotation offset

        R = base * ring_outer
        r = max(8.0, (h[20 + i] % 37) + base * ring_inner * 0.5)
        d = base * (0.06 + 0.22 * (h[30 + i] / 255.0))

        k = (R - r) / r
        total = math.tau * ring_petals
        pts = []
        for j in range(steps + 1):
            t = total * (j / steps) + ring_rot
            x = cx + (R - r) * math.cos(t) + d * math.cos(k * t)
            y = cy + (R - r) * math.sin(t) - d * math.sin(k * t)
            pts.append((x, y))

        d_path = [f"M {pts[0][0]:.3f},{pts[0][1]:.3f}"]
        for (x, y) in pts[1:]:
            d_path.append(f"L {x:.3f},{y:.3f}")
        d_path.append("Z")

        parts.append(
            f"<path d=\"{' '.join(d_path)}\" stroke-width=\"{ring_sw:.3f}\" opacity=\"{ring_opacity:.3f}\"/>"
        )

    parts.append("</g>")
    parts.append("</svg>")
    svg = "\n".join(parts)
    return cairosvg.svg2png(bytestring=svg.encode("utf-8"))


def generate_concentric_circles_png(
    seed: str,
    width: int,
    height: int,
    rings: int = 3,
    ring_spacing: int = 6,
    stroke_color: str = "#28646E",
    stroke_width: float = 1.0,
    opacity: float = 0.24,
) -> bytes:
    """Render elegant concentric circles (perfect rings) as an SVG and return PNG bytes."""
    cx, cy = width / 2.0, height / 2.0
    base = min(width, height)

    h = hashlib.sha256(seed.encode("utf-8")).digest()

    # Seeded slight variation for aesthetics
    rings = max(1, rings)
    ring_spacing = max(1, ring_spacing)
    base_margin = 4 + (h[0] % 3)  # 4..6 px
    max_radius = base / 2.0 - base_margin

    parts = [
        f"<svg width=\"{width}\" height=\"{height}\" viewBox=\"0 0 {width} {height}\" xmlns=\"http://www.w3.org/2000/svg\">",
        f"<g fill=\"none\" stroke=\"{stroke_color}\" stroke-linecap=\"round\" stroke-linejoin=\"round\" opacity=\"{opacity}\">",
    ]

    for i in range(rings):
        r = max_radius - i * ring_spacing
        if r <= 0:
            break
        sw = max(0.5, stroke_width * (1.0 - 0.08 * i))
        parts.append(
            f"<circle cx=\"{cx:.2f}\" cy=\"{cy:.2f}\" r=\"{r:.2f}\" stroke-width=\"{sw:.2f}\"/>"
        )

    parts.append("</g>")
    parts.append("</svg>")
    svg = "\n".join(parts)
    return cairosvg.svg2png(bytestring=svg.encode("utf-8"))


def generate_wavy_rings_png(
    seed: str,
    width: int,
    height: int,
    rings: int = 3,
    ring_spacing: int = 10,
    stroke_color: str = "#28646E",
    stroke_width: float = 1.0,
    opacity: float = 0.24,
    steps: int = 1200,
) -> bytes:
    """Render concentric rings with gentle sinusoidal waves in radius."""
    cx, cy = width / 2.0, height / 2.0
    base = min(width, height)
    h = hashlib.sha256(seed.encode("utf-8")).digest()

    rings = max(1, rings)
    ring_spacing = max(1, ring_spacing)
    base_margin = 180 + (h[0] % 7)  # 4..6 px
    max_radius = base / 2.0 - base_margin

    parts = [
        f"<svg width=\"{width}\" height=\"{height}\" viewBox=\"0 0 {width} {height}\" xmlns=\"http://www.w3.org/2000/svg\">",
        f"<g fill=\"none\" stroke=\"{stroke_color}\" stroke-linecap=\"round\" stroke-linejoin=\"round\" opacity=\"{opacity}\">",
    ]

    for i in range(rings):
        r0 = max_radius - i * ring_spacing
        if r0 <= 0:
            break
        # Seeded wave parameters per ring (modulo hash length for safety)
        L = len(h)
        freq = 120 + (h[(10 + i) % L] % 16)        # lobes
        base_amp = 0.020 + (h[(20 + i) % L] / 255.0) * 0.004
        phase = (h[(30 + i) % L] / 255.0) * math.tau
        sw = max(0.5, stroke_width * (1.0 - 0.1 * i))

        # Per-lobe amplitude table derived from seed
        # amps[k] in ~ [0.6, 1.4] scaling of base_amp
        amps = []
        for k in range(freq):
            b = h[(40 + i * 47 + k) % len(h)] / 255.0
            amps.append(0.6 + 0.8 * b)

        # Low-frequency envelope to vary amplitude smoothly around the ring
        f1 = 2 + (h[(50 + i) % len(h)] % 3)          # 2..4 cycles
        f2 = 3 + (h[(60 + i) % len(h)] % 3)          # 3..5 cycles
        ph1 = (h[(70 + i) % len(h)] / 255.0) * math.tau
        ph2 = (h[(80 + i) % len(h)] / 255.0) * math.tau

        # Build path around circle with radius modulated by sine
        d = []
        for j in range(steps + 1):
            t = (j / steps) * math.tau
            # Determine lobe index and interpolate amplitude between adjacent lobes for smoothness
            u = (t / math.tau) * freq  # lobe-space coordinate
            li = int(u) % freq
            ln = (li + 1) % freq
            frac = u - int(u)
            amp_local = base_amp * (amps[li] * (1.0 - frac) + amps[ln] * frac)
            # Apply smooth envelope (0.7..1.3 scaling)
            env = 1.0 + 0.28 * math.sin(f1 * t + ph1) + 0.18 * math.sin(f2 * t + ph2)
            amp_local *= max(0.5, min(1.6, env))
            r = r0 * (1.0 + amp_local * math.sin(freq * t + phase))
            x = cx + r * math.cos(t)
            y = cy + r * math.sin(t)
            cmd = "M" if j == 0 else "L"
            d.append(f"{cmd} {x:.3f},{y:.3f}")
        d.append("Z")

        parts.append(f"<path d=\"{' '.join(d)}\" stroke-width=\"{sw:.2f}\"/>")

    parts.append("</g>")
    parts.append("</svg>")
    svg = "\n".join(parts)
    return cairosvg.svg2png(bytestring=svg.encode("utf-8"))

