package main

type ColorFormat int

const (
	HEX ColorFormat = iota
	OKLCH
)

type ColorVariant struct {
	Hex   string
	Oklch string
}

type ColorScale map[int]ColorVariant

type ColorName string

const (
	ColorRed     ColorName = "Red"
	ColorOrange  ColorName = "Orange"
	ColorAmber   ColorName = "Amber"
	ColorYellow  ColorName = "Yellow"
	ColorLime    ColorName = "Lime"
	ColorGreen   ColorName = "Green"
	ColorEmerald ColorName = "Emerald"
	ColorTeal    ColorName = "Teal"
	ColorCyan    ColorName = "Cyan"
	ColorSky     ColorName = "Sky"
	ColorBlue    ColorName = "Blue"
	ColorIndigo  ColorName = "Indigo"
	ColorViolet  ColorName = "Violet"
	ColorPurple  ColorName = "Purple"
	ColorFuchsia ColorName = "Fuchsia"
	ColorPink    ColorName = "Pink"
	ColorRose    ColorName = "Rose"
	ColorSlate   ColorName = "Slate"
	ColorGray    ColorName = "Gray"
	ColorZinc    ColorName = "Zinc"
	ColorNeutral ColorName = "Neutral"
	ColorStone   ColorName = "Stone"
)

var Amber = ColorScale{
	50:  {Hex: "#fffbeb", Oklch: "oklch(0.987 0.022 95.277)"},
	100: {Hex: "#fef3c7", Oklch: "oklch(0.962 0.059 95.617)"},
	200: {Hex: "#fde68a", Oklch: "oklch(0.924 0.12 95.746)"},
	300: {Hex: "#fcd34d", Oklch: "oklch(0.879 0.169 91.605)"},
	400: {Hex: "#fbbf24", Oklch: "oklch(0.828 0.189 84.429)"},
	500: {Hex: "#f59e0b", Oklch: "oklch(0.769 0.188 70.08)"},
	600: {Hex: "#d97706", Oklch: "oklch(0.666 0.179 58.318)"},
	700: {Hex: "#b45309", Oklch: "oklch(0.555 0.163 48.998)"},
	800: {Hex: "#92400e", Oklch: "oklch(0.473 0.137 46.201)"},
	900: {Hex: "#78350f", Oklch: "oklch(0.414 0.112 45.904)"},
	950: {Hex: "#451a03", Oklch: "oklch(0.279 0.077 45.635)"},
}

var Blue = ColorScale{
	50:  {Hex: "#eff6ff", Oklch: "Oklch(0.97 0.014 254.604)"},
	100: {Hex: "#dbeafe", Oklch: "Oklch(0.932 0.032 255.585)"},
	200: {Hex: "#bfdbfe", Oklch: "Oklch(0.882 0.059 254.128)"},
	300: {Hex: "#93c5fd", Oklch: "Oklch(0.809 0.105 251.813)"},
	400: {Hex: "#60a5fa", Oklch: "Oklch(0.707 0.165 254.624)"},
	500: {Hex: "#3b82f6", Oklch: "Oklch(0.623 0.214 259.815)"},
	600: {Hex: "#2563eb", Oklch: "Oklch(0.546 0.245 262.881)"},
	700: {Hex: "#1d4ed8", Oklch: "Oklch(0.488 0.243 264.376)"},
	800: {Hex: "#1e40af", Oklch: "Oklch(0.424 0.199 265.638)"},
	900: {Hex: "#1e3a8a", Oklch: "Oklch(0.379 0.146 265.522)"},
	950: {Hex: "#172554", Oklch: "Oklch(0.282 0.091 267.935)"},
}

var Cyan = ColorScale{
	50:  {Hex: "#ecfeff", Oklch: "Oklch(0.984 0.019 200.873)"},
	100: {Hex: "#cffafe", Oklch: "Oklch(0.956 0.045 203.388)"},
	200: {Hex: "#a5f3fc", Oklch: "Oklch(0.917 0.08 205.041)"},
	300: {Hex: "#67e8f9", Oklch: "Oklch(0.865 0.127 207.078)"},
	400: {Hex: "#22d3ee", Oklch: "Oklch(0.789 0.154 211.53)"},
	500: {Hex: "#06b6d4", Oklch: "Oklch(0.715 0.143 215.221)"},
	600: {Hex: "#0891b2", Oklch: "Oklch(0.609 0.126 221.723)"},
	700: {Hex: "#0e7490", Oklch: "Oklch(0.52 0.105 223.128)"},
	800: {Hex: "#155e75", Oklch: "Oklch(0.45 0.085 224.283)"},
	900: {Hex: "#164e63", Oklch: "Oklch(0.398 0.07 227.392)"},
	950: {Hex: "#083344", Oklch: "Oklch(0.302 0.056 229.695)"},
}

var Emerald = ColorScale{
	50:  {Hex: "#ecfdf5", Oklch: "Oklch(0.979 0.021 166.113)"},
	100: {Hex: "#d1fae5", Oklch: "Oklch(0.95 0.052 163.051)"},
	200: {Hex: "#a7f3d0", Oklch: "Oklch(0.905 0.093 164.15)"},
	300: {Hex: "#6ee7b7", Oklch: "Oklch(0.845 0.143 164.978)"},
	400: {Hex: "#34d399", Oklch: "Oklch(0.765 0.177 163.223)"},
	500: {Hex: "#10b981", Oklch: "Oklch(0.696 0.17 162.48)"},
	600: {Hex: "#059669", Oklch: "Oklch(0.596 0.145 163.225)"},
	700: {Hex: "#047857", Oklch: "Oklch(0.508 0.118 165.612)"},
	800: {Hex: "#065f46", Oklch: "Oklch(0.432 0.095 166.913)"},
	900: {Hex: "#064e3b", Oklch: "Oklch(0.378 0.077 168.94)"},
	950: {Hex: "#022c22", Oklch: "Oklch(0.262 0.051 172.552)"},
}

var Fuchsia = ColorScale{
	50:  {Hex: "#fdf4ff", Oklch: "Oklch(0.977 0.017 320.058)"},
	100: {Hex: "#fae8ff", Oklch: "Oklch(0.952 0.037 318.852)"},
	200: {Hex: "#f5d0fe", Oklch: "Oklch(0.903 0.076 319.62)"},
	300: {Hex: "#f0abfc", Oklch: "Oklch(0.833 0.145 321.434)"},
	400: {Hex: "#e879f9", Oklch: "Oklch(0.74 0.238 322.16)"},
	500: {Hex: "#d946ef", Oklch: "Oklch(0.667 0.295 322.15)"},
	600: {Hex: "#c026d3", Oklch: "Oklch(0.591 0.293 322.896)"},
	700: {Hex: "#a21caf", Oklch: "Oklch(0.518 0.253 323.949)"},
	800: {Hex: "#86198f", Oklch: "Oklch(0.452 0.211 324.591)"},
	900: {Hex: "#701a75", Oklch: "Oklch(0.401 0.17 325.612)"},
	950: {Hex: "#4a044e", Oklch: "Oklch(0.293 0.136 325.661)"},
}

var Gray = ColorScale{
	50:  {Hex: "#f9fafb", Oklch: "Oklch(0.985 0.002 247.839)"},
	100: {Hex: "#f3f4f6", Oklch: "Oklch(0.967 0.003 264.542)"},
	200: {Hex: "#e5e7eb", Oklch: "Oklch(0.928 0.006 264.531)"},
	300: {Hex: "#d1d5db", Oklch: "Oklch(0.872 0.01 258.338)"},
	400: {Hex: "#9ca3af", Oklch: "Oklch(0.707 0.022 261.325)"},
	500: {Hex: "#6b7280", Oklch: "Oklch(0.551 0.027 264.364)"},
	600: {Hex: "#4b5563", Oklch: "Oklch(0.446 0.03 256.802)"},
	700: {Hex: "#374151", Oklch: "Oklch(0.373 0.034 259.733)"},
	800: {Hex: "#1f2937", Oklch: "Oklch(0.278 0.033 256.848)"},
	900: {Hex: "#111827", Oklch: "Oklch(0.21 0.034 264.665)"},
	950: {Hex: "#030712", Oklch: "Oklch(0.13 0.028 261.692)"},
}

var Green = ColorScale{
	50:  {Hex: "#f0fdf4", Oklch: "Oklch(0.982 0.018 155.826)"},
	100: {Hex: "#dcfce7", Oklch: "Oklch(0.962 0.044 156.743)"},
	200: {Hex: "#bbf7d0", Oklch: "Oklch(0.925 0.084 155.995)"},
	300: {Hex: "#86efac", Oklch: "Oklch(0.871 0.15 154.449)"},
	400: {Hex: "#4ade80", Oklch: "Oklch(0.792 0.209 151.711)"},
	500: {Hex: "#22c55e", Oklch: "Oklch(0.723 0.219 149.579)"},
	600: {Hex: "#16a34a", Oklch: "Oklch(0.627 0.194 149.214)"},
	700: {Hex: "#15803d", Oklch: "Oklch(0.527 0.154 150.069)"},
	800: {Hex: "#166534", Oklch: "Oklch(0.448 0.119 151.328)"},
	900: {Hex: "#14532d", Oklch: "Oklch(0.393 0.095 152.535)"},
	950: {Hex: "#052e16", Oklch: "Oklch(0.266 0.065 152.934)"},
}

var Indigo = ColorScale{
	50:  {Hex: "#eef2ff", Oklch: "Oklch(0.962 0.018 272.314)"},
	100: {Hex: "#e0e7ff", Oklch: "Oklch(0.93 0.034 272.788)"},
	200: {Hex: "#c7d2fe", Oklch: "Oklch(0.87 0.065 274.039)"},
	300: {Hex: "#a5b4fc", Oklch: "Oklch(0.785 0.115 274.713)"},
	400: {Hex: "#818cf8", Oklch: "Oklch(0.673 0.182 276.935)"},
	500: {Hex: "#6366f1", Oklch: "Oklch(0.585 0.233 277.117)"},
	600: {Hex: "#4f46e5", Oklch: "Oklch(0.511 0.262 276.966)"},
	700: {Hex: "#4338ca", Oklch: "Oklch(0.457 0.24 277.023)"},
	800: {Hex: "#3730a3", Oklch: "Oklch(0.398 0.195 277.366)"},
	900: {Hex: "#312e81", Oklch: "Oklch(0.359 0.144 278.697)"},
	950: {Hex: "#1e1b4b", Oklch: "Oklch(0.257 0.09 281.288)"},
}

var Lime = ColorScale{
	50:  {Hex: "#f7fee7", Oklch: "Oklch(0.986 0.031 120.757)"},
	100: {Hex: "#ecfccb", Oklch: "Oklch(0.967 0.067 122.328)"},
	200: {Hex: "#d9f99d", Oklch: "Oklch(0.938 0.127 124.321)"},
	300: {Hex: "#bef264", Oklch: "Oklch(0.897 0.196 126.665)"},
	400: {Hex: "#a3e635", Oklch: "Oklch(0.841 0.238 128.85)"},
	500: {Hex: "#84cc16", Oklch: "Oklch(0.768 0.233 130.85)"},
	600: {Hex: "#65a30d", Oklch: "Oklch(0.648 0.2 131.684)"},
	700: {Hex: "#4d7c0f", Oklch: "Oklch(0.532 0.157 131.589)"},
	800: {Hex: "#3f6212", Oklch: "Oklch(0.453 0.124 130.933)"},
	900: {Hex: "#365314", Oklch: "Oklch(0.405 0.101 131.063)"},
	950: {Hex: "#1a2e05", Oklch: "Oklch(0.274 0.072 132.109)"},
}

var Neutral = ColorScale{
	50:  {Hex: "#fafafa", Oklch: "Oklch(0.985 0 0)"},
	100: {Hex: "#f5f5f5", Oklch: "Oklch(0.97 0 0)"},
	200: {Hex: "#e5e5e5", Oklch: "Oklch(0.922 0 0)"},
	300: {Hex: "#d4d4d4", Oklch: "Oklch(0.87 0 0)"},
	400: {Hex: "#a3a3a3", Oklch: "Oklch(0.708 0 0)"},
	500: {Hex: "#737373", Oklch: "Oklch(0.556 0 0)"},
	600: {Hex: "#525252", Oklch: "Oklch(0.439 0 0)"},
	700: {Hex: "#404040", Oklch: "Oklch(0.371 0 0)"},
	800: {Hex: "#262626", Oklch: "Oklch(0.269 0 0)"},
	900: {Hex: "#171717", Oklch: "Oklch(0.205 0 0)"},
	950: {Hex: "#0a0a0a", Oklch: "Oklch(0.145 0 0)"},
}

var Orange = ColorScale{
	50:  {Hex: "#fff7ed", Oklch: "Oklch(0.98 0.016 73.684)"},
	100: {Hex: "#ffedd5", Oklch: "Oklch(0.954 0.038 75.164)"},
	200: {Hex: "#fed7aa", Oklch: "Oklch(0.901 0.076 70.697)"},
	300: {Hex: "#fdba74", Oklch: "Oklch(0.837 0.128 66.29)"},
	400: {Hex: "#fb923c", Oklch: "Oklch(0.75 0.183 55.934)"},
	500: {Hex: "#f97316", Oklch: "Oklch(0.705 0.213 47.604)"},
	600: {Hex: "#ea580c", Oklch: "Oklch(0.646 0.222 41.116)"},
	700: {Hex: "#c2410c", Oklch: "Oklch(0.553 0.195 38.402)"},
	800: {Hex: "#9a3412", Oklch: "Oklch(0.47 0.157 37.304)"},
	900: {Hex: "#7c2d12", Oklch: "Oklch(0.408 0.123 38.172)"},
	950: {Hex: "#431407", Oklch: "Oklch(0.266 0.079 36.259)"},
}

var Pink = ColorScale{
	50:  {Hex: "#fdf2f8", Oklch: "Oklch(0.971 0.014 343.198)"},
	100: {Hex: "#fce7f3", Oklch: "Oklch(0.948 0.028 342.258)"},
	200: {Hex: "#fbcfe8", Oklch: "Oklch(0.899 0.061 343.231)"},
	300: {Hex: "#f9a8d4", Oklch: "Oklch(0.823 0.12 346.018)"},
	400: {Hex: "#f472b6", Oklch: "Oklch(0.718 0.202 349.761)"},
	500: {Hex: "#ec4899", Oklch: "Oklch(0.656 0.241 354.308)"},
	600: {Hex: "#db2777", Oklch: "Oklch(0.592 0.249 0.584)"},
	700: {Hex: "#be185d", Oklch: "Oklch(0.525 0.223 3.958)"},
	800: {Hex: "#9d174d", Oklch: "Oklch(0.459 0.187 3.815)"},
	900: {Hex: "#831843", Oklch: "Oklch(0.408 0.153 2.432)"},
	950: {Hex: "#500724", Oklch: "Oklch(0.284 0.109 3.907)"},
}

var Purple = ColorScale{
	50:  {Hex: "#faf5ff", Oklch: "Oklch(0.977 0.014 308.299)"},
	100: {Hex: "#f3e8ff", Oklch: "Oklch(0.946 0.033 307.174)"},
	200: {Hex: "#e9d5ff", Oklch: "Oklch(0.902 0.063 306.703)"},
	300: {Hex: "#d8b4fe", Oklch: "Oklch(0.827 0.119 306.383)"},
	400: {Hex: "#c084fc", Oklch: "Oklch(0.714 0.203 305.504)"},
	500: {Hex: "#a855f7", Oklch: "Oklch(0.627 0.265 303.9)"},
	600: {Hex: "#9333ea", Oklch: "Oklch(0.558 0.288 302.321)"},
	700: {Hex: "#7e22ce", Oklch: "Oklch(0.496 0.265 301.924)"},
	800: {Hex: "#6b21a8", Oklch: "Oklch(0.438 0.218 303.724)"},
	900: {Hex: "#581c87", Oklch: "Oklch(0.381 0.176 304.987)"},
	950: {Hex: "#3b0764", Oklch: "Oklch(0.291 0.149 302.717)"},
}

var Red = ColorScale{
	50:  {Hex: "#fef2f2", Oklch: "Oklch(0.971 0.013 17.38)"},
	100: {Hex: "#fee2e2", Oklch: "Oklch(0.936 0.032 17.717)"},
	200: {Hex: "#fecaca", Oklch: "Oklch(0.885 0.062 18.334)"},
	300: {Hex: "#fca5a5", Oklch: "Oklch(0.808 0.114 19.571)"},
	400: {Hex: "#f87171", Oklch: "Oklch(0.704 0.191 22.216)"},
	500: {Hex: "#ef4444", Oklch: "Oklch(0.637 0.237 25.331)"},
	600: {Hex: "#dc2626", Oklch: "Oklch(0.577 0.245 27.325)"},
	700: {Hex: "#b91c1c", Oklch: "Oklch(0.505 0.213 27.518)"},
	800: {Hex: "#991b1b", Oklch: "Oklch(0.444 0.177 26.899)"},
	900: {Hex: "#7f1d1d", Oklch: "Oklch(0.396 0.141 25.723)"},
	950: {Hex: "#450a0a", Oklch: "Oklch(0.258 0.092 26.042)"},
}

var Rose = ColorScale{
	50:  {Hex: "#fff1f2", Oklch: "Oklch(0.969 0.015 12.422)"},
	100: {Hex: "#ffe4e6", Oklch: "Oklch(0.941 0.03 12.58)"},
	200: {Hex: "#fecdd3", Oklch: "Oklch(0.892 0.058 10.001)"},
	300: {Hex: "#fda4af", Oklch: "Oklch(0.81 0.117 11.638)"},
	400: {Hex: "#fb7185", Oklch: "Oklch(0.712 0.194 13.428)"},
	500: {Hex: "#f43f5e", Oklch: "Oklch(0.645 0.246 16.439)"},
	600: {Hex: "#e11d48", Oklch: "Oklch(0.586 0.253 17.585)"},
	700: {Hex: "#be123c", Oklch: "Oklch(0.514 0.222 16.935)"},
	800: {Hex: "#9f1239", Oklch: "Oklch(0.455 0.188 13.697)"},
	900: {Hex: "#881337", Oklch: "Oklch(0.41 0.159 10.272)"},
	950: {Hex: "#4c0519", Oklch: "Oklch(0.271 0.105 12.094)"},
}

var Sky = ColorScale{
	50:  {Hex: "#f0f9ff", Oklch: "Oklch(0.977 0.013 236.62)"},
	100: {Hex: "#e0f2fe", Oklch: "Oklch(0.951 0.026 236.824)"},
	200: {Hex: "#bae6fd", Oklch: "Oklch(0.901 0.058 230.902)"},
	300: {Hex: "#7dd3fc", Oklch: "Oklch(0.828 0.111 230.318)"},
	400: {Hex: "#38bdf8", Oklch: "Oklch(0.746 0.16 232.661)"},
	500: {Hex: "#0ea5e9", Oklch: "Oklch(0.685 0.169 237.323)"},
	600: {Hex: "#0284c7", Oklch: "Oklch(0.588 0.158 241.966)"},
	700: {Hex: "#0369a1", Oklch: "Oklch(0.5 0.134 242.749)"},
	800: {Hex: "#075985", Oklch: "Oklch(0.443 0.11 240.79)"},
	900: {Hex: "#0c4a6e", Oklch: "Oklch(0.391 0.09 240.876)"},
	950: {Hex: "#082f49", Oklch: "Oklch(0.293 0.066 243.157)"},
}

var Slate = ColorScale{
	50:  {Hex: "#f8fafc", Oklch: "Oklch(0.984 0.003 247.858)"},
	100: {Hex: "#f1f5f9", Oklch: "Oklch(0.968 0.007 247.896)"},
	200: {Hex: "#e2e8f0", Oklch: "Oklch(0.929 0.013 255.508)"},
	300: {Hex: "#cbd5e1", Oklch: "Oklch(0.869 0.022 252.894)"},
	400: {Hex: "#94a3b8", Oklch: "Oklch(0.704 0.04 256.788)"},
	500: {Hex: "#64748b", Oklch: "Oklch(0.554 0.046 257.417)"},
	600: {Hex: "#475569", Oklch: "Oklch(0.446 0.043 257.281)"},
	700: {Hex: "#334155", Oklch: "Oklch(0.372 0.044 257.287)"},
	800: {Hex: "#1e293b", Oklch: "Oklch(0.279 0.041 260.031)"},
	900: {Hex: "#0f172a", Oklch: "Oklch(0.208 0.042 265.755)"},
	950: {Hex: "#020617", Oklch: "Oklch(0.129 0.042 264.695)"},
}

var Stone = ColorScale{
	50:  {Hex: "#fafaf9", Oklch: "Oklch(0.985 0.001 106.423)"},
	100: {Hex: "#f5f5f4", Oklch: "Oklch(0.97 0.001 106.424)"},
	200: {Hex: "#e7e5e4", Oklch: "Oklch(0.923 0.003 48.717)"},
	300: {Hex: "#d6d3d1", Oklch: "Oklch(0.869 0.005 56.366)"},
	400: {Hex: "#a8a29e", Oklch: "Oklch(0.709 0.01 56.259)"},
	500: {Hex: "#78716c", Oklch: "Oklch(0.553 0.013 58.071)"},
	600: {Hex: "#57534e", Oklch: "Oklch(0.444 0.011 73.639)"},
	700: {Hex: "#44403c", Oklch: "Oklch(0.374 0.01 67.558)"},
	800: {Hex: "#292524", Oklch: "Oklch(0.268 0.007 34.298)"},
	900: {Hex: "#1c1917", Oklch: "Oklch(0.216 0.006 56.043)"},
	950: {Hex: "#0c0a09", Oklch: "Oklch(0.147 0.004 49.25)"},
}

var Teal = ColorScale{
	50:  {Hex: "#f0fdfa", Oklch: "Oklch(0.984 0.014 180.72)"},
	100: {Hex: "#ccfbf1", Oklch: "Oklch(0.953 0.051 180.801)"},
	200: {Hex: "#99f6e4", Oklch: "Oklch(0.91 0.096 180.426)"},
	300: {Hex: "#5eead4", Oklch: "Oklch(0.855 0.138 181.071)"},
	400: {Hex: "#2dd4bf", Oklch: "Oklch(0.777 0.152 181.912)"},
	500: {Hex: "#14b8a6", Oklch: "Oklch(0.704 0.14 182.503)"},
	600: {Hex: "#0d9488", Oklch: "Oklch(0.6 0.118 184.704)"},
	700: {Hex: "#0f766e", Oklch: "Oklch(0.511 0.096 186.391)"},
	800: {Hex: "#115e59", Oklch: "Oklch(0.437 0.078 188.216)"},
	900: {Hex: "#134e4a", Oklch: "Oklch(0.386 0.063 188.416)"},
	950: {Hex: "#042f2e", Oklch: "Oklch(0.277 0.046 192.524)"},
}

var Violet = ColorScale{
	50:  {Hex: "#f5f3ff", Oklch: "Oklch(0.969 0.016 293.756)"},
	100: {Hex: "#ede9fe", Oklch: "Oklch(0.943 0.029 294.588)"},
	200: {Hex: "#ddd6fe", Oklch: "Oklch(0.894 0.057 293.283)"},
	300: {Hex: "#c4b5fd", Oklch: "Oklch(0.811 0.111 293.571)"},
	400: {Hex: "#a78bfa", Oklch: "Oklch(0.702 0.183 293.541)"},
	500: {Hex: "#8b5cf6", Oklch: "Oklch(0.606 0.25 292.717)"},
	600: {Hex: "#7c3aed", Oklch: "Oklch(0.541 0.281 293.009)"},
	700: {Hex: "#6d28d9", Oklch: "Oklch(0.491 0.27 292.581)"},
	800: {Hex: "#5b21b6", Oklch: "Oklch(0.432 0.232 292.759)"},
	900: {Hex: "#4c1d95", Oklch: "Oklch(0.38 0.189 293.745)"},
	950: {Hex: "#2e1065", Oklch: "Oklch(0.283 0.141 291.089)"},
}

var Yellow = ColorScale{
	50:  {Hex: "#fefce8", Oklch: "Oklch(0.987 0.026 102.212)"},
	100: {Hex: "#fef9c3", Oklch: "Oklch(0.973 0.071 103.193)"},
	200: {Hex: "#fef08a", Oklch: "Oklch(0.945 0.129 101.54)"},
	300: {Hex: "#fde047", Oklch: "Oklch(0.905 0.182 98.111)"},
	400: {Hex: "#facc15", Oklch: "Oklch(0.852 0.199 91.936)"},
	500: {Hex: "#eab308", Oklch: "Oklch(0.795 0.184 86.047)"},
	600: {Hex: "#ca8a04", Oklch: "Oklch(0.681 0.162 75.834)"},
	700: {Hex: "#a16207", Oklch: "Oklch(0.554 0.135 66.442)"},
	800: {Hex: "#854d0e", Oklch: "Oklch(0.476 0.114 61.907)"},
	900: {Hex: "#713f12", Oklch: "Oklch(0.421 0.095 57.708)"},
	950: {Hex: "#422006", Oklch: "Oklch(0.286 0.066 53.813)"},
}

var Zinc = ColorScale{
	50:  {Hex: "#fafafa", Oklch: "Oklch(0.985 0 0)"},
	100: {Hex: "#f4f4f5", Oklch: "Oklch(0.967 0.001 286.375)"},
	200: {Hex: "#e4e4e7", Oklch: "Oklch(0.92 0.004 286.32)"},
	300: {Hex: "#d4d4d8", Oklch: "Oklch(0.871 0.006 286.286)"},
	400: {Hex: "#a1a1aa", Oklch: "Oklch(0.705 0.015 286.067)"},
	500: {Hex: "#71717a", Oklch: "Oklch(0.552 0.016 285.938)"},
	600: {Hex: "#52525b", Oklch: "Oklch(0.442 0.017 285.786)"},
	700: {Hex: "#3f3f46", Oklch: "Oklch(0.37 0.013 285.805)"},
	800: {Hex: "#27272a", Oklch: "Oklch(0.274 0.006 286.033)"},
	900: {Hex: "#18181b", Oklch: "Oklch(0.21 0.006 285.885)"},
	950: {Hex: "#09090b", Oklch: "Oklch(0.141 0.005 285.823)"},
}

var Palette = map[ColorName]ColorScale{
	ColorRed:     Red,
	ColorOrange:  Orange,
	ColorAmber:   Amber,
	ColorYellow:  Yellow,
	ColorLime:    Lime,
	ColorGreen:   Green,
	ColorEmerald: Emerald,
	ColorTeal:    Teal,
	ColorCyan:    Cyan,
	ColorSky:     Sky,
	ColorBlue:    Blue,
	ColorIndigo:  Indigo,
	ColorViolet:  Violet,
	ColorPurple:  Purple,
	ColorFuchsia: Fuchsia,
	ColorPink:    Pink,
	ColorRose:    Rose,
	ColorSlate:   Slate,
	ColorGray:    Gray,
	ColorZinc:    Zinc,
	ColorNeutral: Neutral,
	ColorStone:   Stone,
}

func (c ColorVariant) GetByFormat(format ColorFormat) string {
	switch format {
	case HEX:
		return c.Hex
	case OKLCH:
		return c.Oklch
	}
	return ""
}
