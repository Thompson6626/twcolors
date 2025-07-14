package main

type ColorType int

const (
	HEX ColorType = iota
	OKLCH
)

type ColorScale map[int]map[ColorType]string

var AMBER = ColorScale{
	50:  {HEX: "#fffbeb", OKLCH: "oklch(0.987 0.022 95.277)"},
	100: {HEX: "#fef3c7", OKLCH: "oklch(0.962 0.059 95.617)"},
	200: {HEX: "#fde68a", OKLCH: "oklch(0.924 0.12 95.746)"},
	300: {HEX: "#fcd34d", OKLCH: "oklch(0.879 0.169 91.605)"},
	400: {HEX: "#fbbf24", OKLCH: "oklch(0.828 0.189 84.429)"},
	500: {HEX: "#f59e0b", OKLCH: "oklch(0.769 0.188 70.08)"},
	600: {HEX: "#d97706", OKLCH: "oklch(0.666 0.179 58.318)"},
	700: {HEX: "#b45309", OKLCH: "oklch(0.555 0.163 48.998)"},
	800: {HEX: "#92400e", OKLCH: "oklch(0.473 0.137 46.201)"},
	900: {HEX: "#78350f", OKLCH: "oklch(0.414 0.112 45.904)"},
	950: {HEX: "#451a03", OKLCH: "oklch(0.279 0.077 45.635)"},
}

var BLUE = ColorScale{
	50:  {HEX: "#eff6ff", OKLCH: "oklch(0.97 0.014 254.604)"},
	100: {HEX: "#dbeafe", OKLCH: "oklch(0.932 0.032 255.585)"},
	200: {HEX: "#bfdbfe", OKLCH: "oklch(0.882 0.059 254.128)"},
	300: {HEX: "#93c5fd", OKLCH: "oklch(0.809 0.105 251.813)"},
	400: {HEX: "#60a5fa", OKLCH: "oklch(0.707 0.165 254.624)"},
	500: {HEX: "#3b82f6", OKLCH: "oklch(0.623 0.214 259.815)"},
	600: {HEX: "#2563eb", OKLCH: "oklch(0.546 0.245 262.881)"},
	700: {HEX: "#1d4ed8", OKLCH: "oklch(0.488 0.243 264.376)"},
	800: {HEX: "#1e40af", OKLCH: "oklch(0.424 0.199 265.638)"},
	900: {HEX: "#1e3a8a", OKLCH: "oklch(0.379 0.146 265.522)"},
	950: {HEX: "#172554", OKLCH: "oklch(0.282 0.091 267.935)"},
}

var CYAN = ColorScale{
	50:  {HEX: "#ecfeff", OKLCH: "oklch(0.984 0.019 200.873)"},
	100: {HEX: "#cffafe", OKLCH: "oklch(0.956 0.045 203.388)"},
	200: {HEX: "#a5f3fc", OKLCH: "oklch(0.917 0.08 205.041)"},
	300: {HEX: "#67e8f9", OKLCH: "oklch(0.865 0.127 207.078)"},
	400: {HEX: "#22d3ee", OKLCH: "oklch(0.789 0.154 211.53)"},
	500: {HEX: "#06b6d4", OKLCH: "oklch(0.715 0.143 215.221)"},
	600: {HEX: "#0891b2", OKLCH: "oklch(0.609 0.126 221.723)"},
	700: {HEX: "#0e7490", OKLCH: "oklch(0.52 0.105 223.128)"},
	800: {HEX: "#155e75", OKLCH: "oklch(0.45 0.085 224.283)"},
	900: {HEX: "#164e63", OKLCH: "oklch(0.398 0.07 227.392)"},
	950: {HEX: "#083344", OKLCH: "oklch(0.302 0.056 229.695)"},
}

var EMERALD = ColorScale{
	50:  {HEX: "#ecfdf5", OKLCH: "oklch(0.979 0.021 166.113)"},
	100: {HEX: "#d1fae5", OKLCH: "oklch(0.95 0.052 163.051)"},
	200: {HEX: "#a7f3d0", OKLCH: "oklch(0.905 0.093 164.15)"},
	300: {HEX: "#6ee7b7", OKLCH: "oklch(0.845 0.143 164.978)"},
	400: {HEX: "#34d399", OKLCH: "oklch(0.765 0.177 163.223)"},
	500: {HEX: "#10b981", OKLCH: "oklch(0.696 0.17 162.48)"},
	600: {HEX: "#059669", OKLCH: "oklch(0.596 0.145 163.225)"},
	700: {HEX: "#047857", OKLCH: "oklch(0.508 0.118 165.612)"},
	800: {HEX: "#065f46", OKLCH: "oklch(0.432 0.095 166.913)"},
	900: {HEX: "#064e3b", OKLCH: "oklch(0.378 0.077 168.94)"},
	950: {HEX: "#022c22", OKLCH: "oklch(0.262 0.051 172.552)"},
}

var FUCHSIA = ColorScale{
	50:  {HEX: "#fdf4ff", OKLCH: "oklch(0.977 0.017 320.058)"},
	100: {HEX: "#fae8ff", OKLCH: "oklch(0.952 0.037 318.852)"},
	200: {HEX: "#f5d0fe", OKLCH: "oklch(0.903 0.076 319.62)"},
	300: {HEX: "#f0abfc", OKLCH: "oklch(0.833 0.145 321.434)"},
	400: {HEX: "#e879f9", OKLCH: "oklch(0.74 0.238 322.16)"},
	500: {HEX: "#d946ef", OKLCH: "oklch(0.667 0.295 322.15)"},
	600: {HEX: "#c026d3", OKLCH: "oklch(0.591 0.293 322.896)"},
	700: {HEX: "#a21caf", OKLCH: "oklch(0.518 0.253 323.949)"},
	800: {HEX: "#86198f", OKLCH: "oklch(0.452 0.211 324.591)"},
	900: {HEX: "#701a75", OKLCH: "oklch(0.401 0.17 325.612)"},
	950: {HEX: "#4a044e", OKLCH: "oklch(0.293 0.136 325.661)"},
}

var GRAY = ColorScale{
	50:  {HEX: "#f9fafb", OKLCH: "oklch(0.985 0.002 247.839)"},
	100: {HEX: "#f3f4f6", OKLCH: "oklch(0.967 0.003 264.542)"},
	200: {HEX: "#e5e7eb", OKLCH: "oklch(0.928 0.006 264.531)"},
	300: {HEX: "#d1d5db", OKLCH: "oklch(0.872 0.01 258.338)"},
	400: {HEX: "#9ca3af", OKLCH: "oklch(0.707 0.022 261.325)"},
	500: {HEX: "#6b7280", OKLCH: "oklch(0.551 0.027 264.364)"},
	600: {HEX: "#4b5563", OKLCH: "oklch(0.446 0.03 256.802)"},
	700: {HEX: "#374151", OKLCH: "oklch(0.373 0.034 259.733)"},
	800: {HEX: "#1f2937", OKLCH: "oklch(0.278 0.033 256.848)"},
	900: {HEX: "#111827", OKLCH: "oklch(0.21 0.034 264.665)"},
	950: {HEX: "#030712", OKLCH: "oklch(0.13 0.028 261.692)"},
}

var GREEN = ColorScale{
	50:  {HEX: "#f0fdf4", OKLCH: "oklch(0.982 0.018 155.826)"},
	100: {HEX: "#dcfce7", OKLCH: "oklch(0.962 0.044 156.743)"},
	200: {HEX: "#bbf7d0", OKLCH: "oklch(0.925 0.084 155.995)"},
	300: {HEX: "#86efac", OKLCH: "oklch(0.871 0.15 154.449)"},
	400: {HEX: "#4ade80", OKLCH: "oklch(0.792 0.209 151.711)"},
	500: {HEX: "#22c55e", OKLCH: "oklch(0.723 0.219 149.579)"},
	600: {HEX: "#16a34a", OKLCH: "oklch(0.627 0.194 149.214)"},
	700: {HEX: "#15803d", OKLCH: "oklch(0.527 0.154 150.069)"},
	800: {HEX: "#166534", OKLCH: "oklch(0.448 0.119 151.328)"},
	900: {HEX: "#14532d", OKLCH: "oklch(0.393 0.095 152.535)"},
	950: {HEX: "#052e16", OKLCH: "oklch(0.266 0.065 152.934)"},
}

var INDIGO = ColorScale{
	50:  {HEX: "#eef2ff", OKLCH: "oklch(0.962 0.018 272.314)"},
	100: {HEX: "#e0e7ff", OKLCH: "oklch(0.93 0.034 272.788)"},
	200: {HEX: "#c7d2fe", OKLCH: "oklch(0.87 0.065 274.039)"},
	300: {HEX: "#a5b4fc", OKLCH: "oklch(0.785 0.115 274.713)"},
	400: {HEX: "#818cf8", OKLCH: "oklch(0.673 0.182 276.935)"},
	500: {HEX: "#6366f1", OKLCH: "oklch(0.585 0.233 277.117)"},
	600: {HEX: "#4f46e5", OKLCH: "oklch(0.511 0.262 276.966)"},
	700: {HEX: "#4338ca", OKLCH: "oklch(0.457 0.24 277.023)"},
	800: {HEX: "#3730a3", OKLCH: "oklch(0.398 0.195 277.366)"},
	900: {HEX: "#312e81", OKLCH: "oklch(0.359 0.144 278.697)"},
	950: {HEX: "#1e1b4b", OKLCH: "oklch(0.257 0.09 281.288)"},
}

var LIME = ColorScale{
	50:  {HEX: "#f7fee7", OKLCH: "oklch(0.986 0.031 120.757)"},
	100: {HEX: "#ecfccb", OKLCH: "oklch(0.967 0.067 122.328)"},
	200: {HEX: "#d9f99d", OKLCH: "oklch(0.938 0.127 124.321)"},
	300: {HEX: "#bef264", OKLCH: "oklch(0.897 0.196 126.665)"},
	400: {HEX: "#a3e635", OKLCH: "oklch(0.841 0.238 128.85)"},
	500: {HEX: "#84cc16", OKLCH: "oklch(0.768 0.233 130.85)"},
	600: {HEX: "#65a30d", OKLCH: "oklch(0.648 0.2 131.684)"},
	700: {HEX: "#4d7c0f", OKLCH: "oklch(0.532 0.157 131.589)"},
	800: {HEX: "#3f6212", OKLCH: "oklch(0.453 0.124 130.933)"},
	900: {HEX: "#365314", OKLCH: "oklch(0.405 0.101 131.063)"},
	950: {HEX: "#1a2e05", OKLCH: "oklch(0.274 0.072 132.109)"},
}

var NEUTRAL = ColorScale{
	50:  {HEX: "#fafafa", OKLCH: "oklch(0.985 0 0)"},
	100: {HEX: "#f5f5f5", OKLCH: "oklch(0.97 0 0)"},
	200: {HEX: "#e5e5e5", OKLCH: "oklch(0.922 0 0)"},
	300: {HEX: "#d4d4d4", OKLCH: "oklch(0.87 0 0)"},
	400: {HEX: "#a3a3a3", OKLCH: "oklch(0.708 0 0)"},
	500: {HEX: "#737373", OKLCH: "oklch(0.556 0 0)"},
	600: {HEX: "#525252", OKLCH: "oklch(0.439 0 0)"},
	700: {HEX: "#404040", OKLCH: "oklch(0.371 0 0)"},
	800: {HEX: "#262626", OKLCH: "oklch(0.269 0 0)"},
	900: {HEX: "#171717", OKLCH: "oklch(0.205 0 0)"},
	950: {HEX: "#0a0a0a", OKLCH: "oklch(0.145 0 0)"},
}

var ORANGE = ColorScale{
	50:  {HEX: "#fff7ed", OKLCH: "oklch(0.98 0.016 73.684)"},
	100: {HEX: "#ffedd5", OKLCH: "oklch(0.954 0.038 75.164)"},
	200: {HEX: "#fed7aa", OKLCH: "oklch(0.901 0.076 70.697)"},
	300: {HEX: "#fdba74", OKLCH: "oklch(0.837 0.128 66.29)"},
	400: {HEX: "#fb923c", OKLCH: "oklch(0.75 0.183 55.934)"},
	500: {HEX: "#f97316", OKLCH: "oklch(0.705 0.213 47.604)"},
	600: {HEX: "#ea580c", OKLCH: "oklch(0.646 0.222 41.116)"},
	700: {HEX: "#c2410c", OKLCH: "oklch(0.553 0.195 38.402)"},
	800: {HEX: "#9a3412", OKLCH: "oklch(0.47 0.157 37.304)"},
	900: {HEX: "#7c2d12", OKLCH: "oklch(0.408 0.123 38.172)"},
	950: {HEX: "#431407", OKLCH: "oklch(0.266 0.079 36.259)"},
}

var PINK = ColorScale{
	50:  {HEX: "#fdf2f8", OKLCH: "oklch(0.971 0.014 343.198)"},
	100: {HEX: "#fce7f3", OKLCH: "oklch(0.948 0.028 342.258)"},
	200: {HEX: "#fbcfe8", OKLCH: "oklch(0.899 0.061 343.231)"},
	300: {HEX: "#f9a8d4", OKLCH: "oklch(0.823 0.12 346.018)"},
	400: {HEX: "#f472b6", OKLCH: "oklch(0.718 0.202 349.761)"},
	500: {HEX: "#ec4899", OKLCH: "oklch(0.656 0.241 354.308)"},
	600: {HEX: "#db2777", OKLCH: "oklch(0.592 0.249 0.584)"},
	700: {HEX: "#be185d", OKLCH: "oklch(0.525 0.223 3.958)"},
	800: {HEX: "#9d174d", OKLCH: "oklch(0.459 0.187 3.815)"},
	900: {HEX: "#831843", OKLCH: "oklch(0.408 0.153 2.432)"},
	950: {HEX: "#500724", OKLCH: "oklch(0.284 0.109 3.907)"},
}

var PURPLE = ColorScale{
	50:  {HEX: "#faf5ff", OKLCH: "oklch(0.977 0.014 308.299)"},
	100: {HEX: "#f3e8ff", OKLCH: "oklch(0.946 0.033 307.174)"},
	200: {HEX: "#e9d5ff", OKLCH: "oklch(0.902 0.063 306.703)"},
	300: {HEX: "#d8b4fe", OKLCH: "oklch(0.827 0.119 306.383)"},
	400: {HEX: "#c084fc", OKLCH: "oklch(0.714 0.203 305.504)"},
	500: {HEX: "#a855f7", OKLCH: "oklch(0.627 0.265 303.9)"},
	600: {HEX: "#9333ea", OKLCH: "oklch(0.558 0.288 302.321)"},
	700: {HEX: "#7e22ce", OKLCH: "oklch(0.496 0.265 301.924)"},
	800: {HEX: "#6b21a8", OKLCH: "oklch(0.438 0.218 303.724)"},
	900: {HEX: "#581c87", OKLCH: "oklch(0.381 0.176 304.987)"},
	950: {HEX: "#3b0764", OKLCH: "oklch(0.291 0.149 302.717)"},
}

var RED = ColorScale{
	50:  {HEX: "#fef2f2", OKLCH: "oklch(0.971 0.013 17.38)"},
	100: {HEX: "#fee2e2", OKLCH: "oklch(0.936 0.032 17.717)"},
	200: {HEX: "#fecaca", OKLCH: "oklch(0.885 0.062 18.334)"},
	300: {HEX: "#fca5a5", OKLCH: "oklch(0.808 0.114 19.571)"},
	400: {HEX: "#f87171", OKLCH: "oklch(0.704 0.191 22.216)"},
	500: {HEX: "#ef4444", OKLCH: "oklch(0.637 0.237 25.331)"},
	600: {HEX: "#dc2626", OKLCH: "oklch(0.577 0.245 27.325)"},
	700: {HEX: "#b91c1c", OKLCH: "oklch(0.505 0.213 27.518)"},
	800: {HEX: "#991b1b", OKLCH: "oklch(0.444 0.177 26.899)"},
	900: {HEX: "#7f1d1d", OKLCH: "oklch(0.396 0.141 25.723)"},
	950: {HEX: "#450a0a", OKLCH: "oklch(0.258 0.092 26.042)"},
}

var ROSE = ColorScale{
	50:  {HEX: "#fff1f2", OKLCH: "oklch(0.969 0.015 12.422)"},
	100: {HEX: "#ffe4e6", OKLCH: "oklch(0.941 0.03 12.58)"},
	200: {HEX: "#fecdd3", OKLCH: "oklch(0.892 0.058 10.001)"},
	300: {HEX: "#fda4af", OKLCH: "oklch(0.81 0.117 11.638)"},
	400: {HEX: "#fb7185", OKLCH: "oklch(0.712 0.194 13.428)"},
	500: {HEX: "#f43f5e", OKLCH: "oklch(0.645 0.246 16.439)"},
	600: {HEX: "#e11d48", OKLCH: "oklch(0.586 0.253 17.585)"},
	700: {HEX: "#be123c", OKLCH: "oklch(0.514 0.222 16.935)"},
	800: {HEX: "#9f1239", OKLCH: "oklch(0.455 0.188 13.697)"},
	900: {HEX: "#881337", OKLCH: "oklch(0.41 0.159 10.272)"},
	950: {HEX: "#4c0519", OKLCH: "oklch(0.271 0.105 12.094)"},
}

var SKY = ColorScale{
	50:  {HEX: "#f0f9ff", OKLCH: "oklch(0.977 0.013 236.62)"},
	100: {HEX: "#e0f2fe", OKLCH: "oklch(0.951 0.026 236.824)"},
	200: {HEX: "#bae6fd", OKLCH: "oklch(0.901 0.058 230.902)"},
	300: {HEX: "#7dd3fc", OKLCH: "oklch(0.828 0.111 230.318)"},
	400: {HEX: "#38bdf8", OKLCH: "oklch(0.746 0.16 232.661)"},
	500: {HEX: "#0ea5e9", OKLCH: "oklch(0.685 0.169 237.323)"},
	600: {HEX: "#0284c7", OKLCH: "oklch(0.588 0.158 241.966)"},
	700: {HEX: "#0369a1", OKLCH: "oklch(0.5 0.134 242.749)"},
	800: {HEX: "#075985", OKLCH: "oklch(0.443 0.11 240.79)"},
	900: {HEX: "#0c4a6e", OKLCH: "oklch(0.391 0.09 240.876)"},
	950: {HEX: "#082f49", OKLCH: "oklch(0.293 0.066 243.157)"},
}

var SLATE = ColorScale{
	50:  {HEX: "#f8fafc", OKLCH: "oklch(0.984 0.003 247.858)"},
	100: {HEX: "#f1f5f9", OKLCH: "oklch(0.968 0.007 247.896)"},
	200: {HEX: "#e2e8f0", OKLCH: "oklch(0.929 0.013 255.508)"},
	300: {HEX: "#cbd5e1", OKLCH: "oklch(0.869 0.022 252.894)"},
	400: {HEX: "#94a3b8", OKLCH: "oklch(0.704 0.04 256.788)"},
	500: {HEX: "#64748b", OKLCH: "oklch(0.554 0.046 257.417)"},
	600: {HEX: "#475569", OKLCH: "oklch(0.446 0.043 257.281)"},
	700: {HEX: "#334155", OKLCH: "oklch(0.372 0.044 257.287)"},
	800: {HEX: "#1e293b", OKLCH: "oklch(0.279 0.041 260.031)"},
	900: {HEX: "#0f172a", OKLCH: "oklch(0.208 0.042 265.755)"},
	950: {HEX: "#020617", OKLCH: "oklch(0.129 0.042 264.695)"},
}

var STONE = ColorScale{
	50:  {HEX: "#fafaf9", OKLCH: "oklch(0.985 0.001 106.423)"},
	100: {HEX: "#f5f5f4", OKLCH: "oklch(0.97 0.001 106.424)"},
	200: {HEX: "#e7e5e4", OKLCH: "oklch(0.923 0.003 48.717)"},
	300: {HEX: "#d6d3d1", OKLCH: "oklch(0.869 0.005 56.366)"},
	400: {HEX: "#a8a29e", OKLCH: "oklch(0.709 0.01 56.259)"},
	500: {HEX: "#78716c", OKLCH: "oklch(0.553 0.013 58.071)"},
	600: {HEX: "#57534e", OKLCH: "oklch(0.444 0.011 73.639)"},
	700: {HEX: "#44403c", OKLCH: "oklch(0.374 0.01 67.558)"},
	800: {HEX: "#292524", OKLCH: "oklch(0.268 0.007 34.298)"},
	900: {HEX: "#1c1917", OKLCH: "oklch(0.216 0.006 56.043)"},
	950: {HEX: "#0c0a09", OKLCH: "oklch(0.147 0.004 49.25)"},
}

var TEAL = ColorScale{
	50:  {HEX: "#f0fdfa", OKLCH: "oklch(0.984 0.014 180.72)"},
	100: {HEX: "#ccfbf1", OKLCH: "oklch(0.953 0.051 180.801)"},
	200: {HEX: "#99f6e4", OKLCH: "oklch(0.91 0.096 180.426)"},
	300: {HEX: "#5eead4", OKLCH: "oklch(0.855 0.138 181.071)"},
	400: {HEX: "#2dd4bf", OKLCH: "oklch(0.777 0.152 181.912)"},
	500: {HEX: "#14b8a6", OKLCH: "oklch(0.704 0.14 182.503)"},
	600: {HEX: "#0d9488", OKLCH: "oklch(0.6 0.118 184.704)"},
	700: {HEX: "#0f766e", OKLCH: "oklch(0.511 0.096 186.391)"},
	800: {HEX: "#115e59", OKLCH: "oklch(0.437 0.078 188.216)"},
	900: {HEX: "#134e4a", OKLCH: "oklch(0.386 0.063 188.416)"},
	950: {HEX: "#042f2e", OKLCH: "oklch(0.277 0.046 192.524)"},
}

var VIOLET = ColorScale{
	50:  {HEX: "#f5f3ff", OKLCH: "oklch(0.969 0.016 293.756)"},
	100: {HEX: "#ede9fe", OKLCH: "oklch(0.943 0.029 294.588)"},
	200: {HEX: "#ddd6fe", OKLCH: "oklch(0.894 0.057 293.283)"},
	300: {HEX: "#c4b5fd", OKLCH: "oklch(0.811 0.111 293.571)"},
	400: {HEX: "#a78bfa", OKLCH: "oklch(0.702 0.183 293.541)"},
	500: {HEX: "#8b5cf6", OKLCH: "oklch(0.606 0.25 292.717)"},
	600: {HEX: "#7c3aed", OKLCH: "oklch(0.541 0.281 293.009)"},
	700: {HEX: "#6d28d9", OKLCH: "oklch(0.491 0.27 292.581)"},
	800: {HEX: "#5b21b6", OKLCH: "oklch(0.432 0.232 292.759)"},
	900: {HEX: "#4c1d95", OKLCH: "oklch(0.38 0.189 293.745)"},
	950: {HEX: "#2e1065", OKLCH: "oklch(0.283 0.141 291.089)"},
}

var YELLOW = ColorScale{
	50:  {HEX: "#fefce8", OKLCH: "oklch(0.987 0.026 102.212)"},
	100: {HEX: "#fef9c3", OKLCH: "oklch(0.973 0.071 103.193)"},
	200: {HEX: "#fef08a", OKLCH: "oklch(0.945 0.129 101.54)"},
	300: {HEX: "#fde047", OKLCH: "oklch(0.905 0.182 98.111)"},
	400: {HEX: "#facc15", OKLCH: "oklch(0.852 0.199 91.936)"},
	500: {HEX: "#eab308", OKLCH: "oklch(0.795 0.184 86.047)"},
	600: {HEX: "#ca8a04", OKLCH: "oklch(0.681 0.162 75.834)"},
	700: {HEX: "#a16207", OKLCH: "oklch(0.554 0.135 66.442)"},
	800: {HEX: "#854d0e", OKLCH: "oklch(0.476 0.114 61.907)"},
	900: {HEX: "#713f12", OKLCH: "oklch(0.421 0.095 57.708)"},
	950: {HEX: "#422006", OKLCH: "oklch(0.286 0.066 53.813)"},
}

var ZINC = ColorScale{
	50:  {HEX: "#fafafa", OKLCH: "oklch(0.985 0 0)"},
	100: {HEX: "#f4f4f5", OKLCH: "oklch(0.967 0.001 286.375)"},
	200: {HEX: "#e4e4e7", OKLCH: "oklch(0.92 0.004 286.32)"},
	300: {HEX: "#d4d4d8", OKLCH: "oklch(0.871 0.006 286.286)"},
	400: {HEX: "#a1a1aa", OKLCH: "oklch(0.705 0.015 286.067)"},
	500: {HEX: "#71717a", OKLCH: "oklch(0.552 0.016 285.938)"},
	600: {HEX: "#52525b", OKLCH: "oklch(0.442 0.017 285.786)"},
	700: {HEX: "#3f3f46", OKLCH: "oklch(0.37 0.013 285.805)"},
	800: {HEX: "#27272a", OKLCH: "oklch(0.274 0.006 286.033)"},
	900: {HEX: "#18181b", OKLCH: "oklch(0.21 0.006 285.885)"},
	950: {HEX: "#09090b", OKLCH: "oklch(0.141 0.005 285.823)"},
}

var Palette = map[string]ColorScale{
	"RED":     RED,
	"ORANGE":  ORANGE,
	"AMBER":   AMBER,
	"YELLOW":  YELLOW,
	"LIME":    LIME,
	"GREEN":   GREEN,
	"EMERALD": EMERALD,
	"TEAL":    TEAL,
	"CYAN":    CYAN,
	"SKY":     SKY,
	"BLUE":    BLUE,
	"INDIGO":  INDIGO,
	"VIOLET":  VIOLET,
	"PURPLE":  PURPLE,
	"FUCHSIA": FUCHSIA,
	"PINK":    PINK,
	"ROSE":    ROSE,
	"SLATE":   SLATE,
	"GRAY":    GRAY,
	"ZINC":    ZINC,
	"NEUTRAL": NEUTRAL,
	"STONE":   STONE,
}
