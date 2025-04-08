const unicodeMapping = {
  "circle-1": "\uE0EE",
  "circle-2": "\uE0EF",
  "circle-3": "\uE0F0",
  "circle-4": "\uE0F1",
  "circle-5": "\uE0F2",
  "circle-check": "\uF058",
  "circle-exclamation": "\uF06A",
  "link-simple": "\uE1CD",
  "triangle-exclamation": "\uF071",
} as const;

export default unicodeMapping;

export type SupportedIcon = keyof typeof unicodeMapping;
