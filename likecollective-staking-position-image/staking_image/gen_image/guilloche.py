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

