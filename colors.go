package main

import "github.com/charmbracelet/lipgloss"

type ColorScale map[int]lipgloss.Color

var WHITE = lipgloss.Color("#000")

var AMBER = ColorScale{
	50:  lipgloss.Color("#fffbeb"),
	100: lipgloss.Color("#fef3c7"),
	200: lipgloss.Color("#fde68a"),
	300: lipgloss.Color("#fcd34d"),
	400: lipgloss.Color("#fbbf24"),
	500: lipgloss.Color("#f59e0b"),
	600: lipgloss.Color("#d97706"),
	700: lipgloss.Color("#b45309"),
	800: lipgloss.Color("#92400e"),
	900: lipgloss.Color("#78350f"),
	950: lipgloss.Color("#451a03"),
}

var BLUE = ColorScale{
	50:  lipgloss.Color("#eff6ff"),
	100: lipgloss.Color("#dbeafe"),
	200: lipgloss.Color("#bfdbfe"),
	300: lipgloss.Color("#93c5fd"),
	400: lipgloss.Color("#60a5fa"),
	500: lipgloss.Color("#3b82f6"),
	600: lipgloss.Color("#2563eb"),
	700: lipgloss.Color("#1d4ed8"),
	800: lipgloss.Color("#1e40af"),
	900: lipgloss.Color("#1e3a8a"),
	950: lipgloss.Color("#172554"),
}

var CYAN = ColorScale{
	50:  lipgloss.Color("#ecfeff"),
	100: lipgloss.Color("#cffafe"),
	200: lipgloss.Color("#a5f3fc"),
	300: lipgloss.Color("#67e8f9"),
	400: lipgloss.Color("#22d3ee"),
	500: lipgloss.Color("#06b6d4"),
	600: lipgloss.Color("#0891b2"),
	700: lipgloss.Color("#0e7490"),
	800: lipgloss.Color("#155e75"),
	900: lipgloss.Color("#164e63"),
	950: lipgloss.Color("#083344"),
}

var EMERALD = ColorScale{
	50:  lipgloss.Color("#ecfdf5"),
	100: lipgloss.Color("#d1fae5"),
	200: lipgloss.Color("#a7f3d0"),
	300: lipgloss.Color("#6ee7b7"),
	400: lipgloss.Color("#34d399"),
	500: lipgloss.Color("#10b981"),
	600: lipgloss.Color("#059669"),
	700: lipgloss.Color("#047857"),
	800: lipgloss.Color("#065f46"),
	900: lipgloss.Color("#064e3b"),
	950: lipgloss.Color("#022c22"),
}

var FUCHSIA = ColorScale{
	50:  lipgloss.Color("#fdf4ff"),
	100: lipgloss.Color("#fae8ff"),
	200: lipgloss.Color("#f5d0fe"),
	300: lipgloss.Color("#f0abfc"),
	400: lipgloss.Color("#e879f9"),
	500: lipgloss.Color("#d946ef"),
	600: lipgloss.Color("#c026d3"),
	700: lipgloss.Color("#a21caf"),
	800: lipgloss.Color("#86198f"),
	900: lipgloss.Color("#701a75"),
	950: lipgloss.Color("#4a044e"),
}

var GRAY = ColorScale{
	50:  lipgloss.Color("#f9fafb"),
	100: lipgloss.Color("#f3f4f6"),
	200: lipgloss.Color("#e5e7eb"),
	300: lipgloss.Color("#d1d5db"),
	400: lipgloss.Color("#9ca3af"),
	500: lipgloss.Color("#6b7280"),
	600: lipgloss.Color("#4b5563"),
	700: lipgloss.Color("#374151"),
	800: lipgloss.Color("#1f2937"),
	900: lipgloss.Color("#111827"),
	950: lipgloss.Color("#030712"),
}

var GREEN = ColorScale{
	50:  lipgloss.Color("#f0fdf4"),
	100: lipgloss.Color("#dcfce7"),
	200: lipgloss.Color("#bbf7d0"),
	300: lipgloss.Color("#86efac"),
	400: lipgloss.Color("#4ade80"),
	500: lipgloss.Color("#22c55e"),
	600: lipgloss.Color("#16a34a"),
	700: lipgloss.Color("#15803d"),
	800: lipgloss.Color("#166534"),
	900: lipgloss.Color("#14532d"),
	950: lipgloss.Color("#052e16"),
}

var INDIGO = ColorScale{
	50:  lipgloss.Color("#eef2ff"),
	100: lipgloss.Color("#e0e7ff"),
	200: lipgloss.Color("#c7d2fe"),
	300: lipgloss.Color("#a5b4fc"),
	400: lipgloss.Color("#818cf8"),
	500: lipgloss.Color("#6366f1"),
	600: lipgloss.Color("#4f46e5"),
	700: lipgloss.Color("#4338ca"),
	800: lipgloss.Color("#3730a3"),
	900: lipgloss.Color("#312e81"),
	950: lipgloss.Color("#1e1b4b"),
}

var LIME = ColorScale{
	50:  lipgloss.Color("#f7fee7"),
	100: lipgloss.Color("#ecfccb"),
	200: lipgloss.Color("#d9f99d"),
	300: lipgloss.Color("#bef264"),
	400: lipgloss.Color("#a3e635"),
	500: lipgloss.Color("#84cc16"),
	600: lipgloss.Color("#65a30d"),
	700: lipgloss.Color("#4d7c0f"),
	800: lipgloss.Color("#3f6212"),
	900: lipgloss.Color("#365314"),
	950: lipgloss.Color("#1a2e05"),
}

var NEUTRAL = ColorScale{
	50:  lipgloss.Color("#fafafa"),
	100: lipgloss.Color("#f5f5f5"),
	200: lipgloss.Color("#e5e5e5"),
	300: lipgloss.Color("#d4d4d4"),
	400: lipgloss.Color("#a3a3a3"),
	500: lipgloss.Color("#737373"),
	600: lipgloss.Color("#525252"),
	700: lipgloss.Color("#404040"),
	800: lipgloss.Color("#262626"),
	900: lipgloss.Color("#171717"),
	950: lipgloss.Color("#0a0a0a"),
}

var ORANGE = ColorScale{
	50:  lipgloss.Color("#fff7ed"),
	100: lipgloss.Color("#ffedd5"),
	200: lipgloss.Color("#fed7aa"),
	300: lipgloss.Color("#fdba74"),
	400: lipgloss.Color("#fb923c"),
	500: lipgloss.Color("#f97316"),
	600: lipgloss.Color("#ea580c"),
	700: lipgloss.Color("#c2410c"),
	800: lipgloss.Color("#9a3412"),
	900: lipgloss.Color("#7c2d12"),
	950: lipgloss.Color("#431407"),
}

var PINK = ColorScale{
	50:  lipgloss.Color("#fdf2f8"),
	100: lipgloss.Color("#fce7f3"),
	200: lipgloss.Color("#fbcfe8"),
	300: lipgloss.Color("#f9a8d4"),
	400: lipgloss.Color("#f472b6"),
	500: lipgloss.Color("#ec4899"),
	600: lipgloss.Color("#db2777"),
	700: lipgloss.Color("#be185d"),
	800: lipgloss.Color("#9d174d"),
	900: lipgloss.Color("#831843"),
	950: lipgloss.Color("#500724"),
}

var PURPLE = ColorScale{
	50:  lipgloss.Color("#faf5ff"),
	100: lipgloss.Color("#f3e8ff"),
	200: lipgloss.Color("#e9d5ff"),
	300: lipgloss.Color("#d8b4fe"),
	400: lipgloss.Color("#c084fc"),
	500: lipgloss.Color("#a855f7"),
	600: lipgloss.Color("#9333ea"),
	700: lipgloss.Color("#7e22ce"),
	800: lipgloss.Color("#6b21a8"),
	900: lipgloss.Color("#581c87"),
	950: lipgloss.Color("#3b0764"),
}

var RED = ColorScale{
	50:  lipgloss.Color("#fef2f2"),
	100: lipgloss.Color("#fee2e2"),
	200: lipgloss.Color("#fecaca"),
	300: lipgloss.Color("#fca5a5"),
	400: lipgloss.Color("#f87171"),
	500: lipgloss.Color("#ef4444"),
	600: lipgloss.Color("#dc2626"),
	700: lipgloss.Color("#b91c1c"),
	800: lipgloss.Color("#991b1b"),
	900: lipgloss.Color("#7f1d1d"),
	950: lipgloss.Color("#450a0a"),
}

var ROSE = ColorScale{
	50:  lipgloss.Color("#fff1f2"),
	100: lipgloss.Color("#ffe4e6"),
	200: lipgloss.Color("#fecdd3"),
	300: lipgloss.Color("#fda4af"),
	400: lipgloss.Color("#fb7185"),
	500: lipgloss.Color("#f43f5e"),
	600: lipgloss.Color("#e11d48"),
	700: lipgloss.Color("#be123c"),
	800: lipgloss.Color("#9f1239"),
	900: lipgloss.Color("#881337"),
	950: lipgloss.Color("#4c0519"),
}

var SKY = ColorScale{
	50:  lipgloss.Color("#f0f9ff"),
	100: lipgloss.Color("#e0f2fe"),
	200: lipgloss.Color("#bae6fd"),
	300: lipgloss.Color("#7dd3fc"),
	400: lipgloss.Color("#38bdf8"),
	500: lipgloss.Color("#0ea5e9"),
	600: lipgloss.Color("#0284c7"),
	700: lipgloss.Color("#0369a1"),
	800: lipgloss.Color("#075985"),
	900: lipgloss.Color("#0c4a6e"),
	950: lipgloss.Color("#082f49"),
}

var SLATE = ColorScale{
	50:  lipgloss.Color("#f8fafc"),
	100: lipgloss.Color("#f1f5f9"),
	200: lipgloss.Color("#e2e8f0"),
	300: lipgloss.Color("#cbd5e1"),
	400: lipgloss.Color("#94a3b8"),
	500: lipgloss.Color("#64748b"),
	600: lipgloss.Color("#475569"),
	700: lipgloss.Color("#334155"),
	800: lipgloss.Color("#1e293b"),
	900: lipgloss.Color("#0f172a"),
	950: lipgloss.Color("#020617"),
}

var STONE = ColorScale{
	50:  lipgloss.Color("#fafaf9"),
	100: lipgloss.Color("#f5f5f4"),
	200: lipgloss.Color("#e7e5e4"),
	300: lipgloss.Color("#d6d3d1"),
	400: lipgloss.Color("#a8a29e"),
	500: lipgloss.Color("#78716c"),
	600: lipgloss.Color("#57534e"),
	700: lipgloss.Color("#44403c"),
	800: lipgloss.Color("#292524"),
	900: lipgloss.Color("#1c1917"),
	950: lipgloss.Color("#0c0a09"),
}

var TEAL = ColorScale{
	50:  lipgloss.Color("#f0fdfa"),
	100: lipgloss.Color("#ccfbf1"),
	200: lipgloss.Color("#99f6e4"),
	300: lipgloss.Color("#5eead4"),
	400: lipgloss.Color("#2dd4bf"),
	500: lipgloss.Color("#14b8a6"),
	600: lipgloss.Color("#0d9488"),
	700: lipgloss.Color("#0f766e"),
	800: lipgloss.Color("#115e59"),
	900: lipgloss.Color("#134e4a"),
	950: lipgloss.Color("#042f2e"),
}

var VIOLET = ColorScale{
	50:  lipgloss.Color("#f5f3ff"),
	100: lipgloss.Color("#ede9fe"),
	200: lipgloss.Color("#ddd6fe"),
	300: lipgloss.Color("#c4b5fd"),
	400: lipgloss.Color("#a78bfa"),
	500: lipgloss.Color("#8b5cf6"),
	600: lipgloss.Color("#7c3aed"),
	700: lipgloss.Color("#6d28d9"),
	800: lipgloss.Color("#5b21b6"),
	900: lipgloss.Color("#4c1d95"),
	950: lipgloss.Color("#2e1065"),
}

var YELLOW = ColorScale{
	50:  lipgloss.Color("#fefce8"),
	100: lipgloss.Color("#fef9c3"),
	200: lipgloss.Color("#fef08a"),
	300: lipgloss.Color("#fde047"),
	400: lipgloss.Color("#facc15"),
	500: lipgloss.Color("#eab308"),
	600: lipgloss.Color("#ca8a04"),
	700: lipgloss.Color("#a16207"),
	800: lipgloss.Color("#854d0e"),
	900: lipgloss.Color("#713f12"),
	950: lipgloss.Color("#422006"),
}

var ZINC = ColorScale{
	50:  lipgloss.Color("#fafafa"),
	100: lipgloss.Color("#f4f4f5"),
	200: lipgloss.Color("#e4e4e7"),
	300: lipgloss.Color("#d4d4d8"),
	400: lipgloss.Color("#a1a1aa"),
	500: lipgloss.Color("#71717a"),
	600: lipgloss.Color("#52525b"),
	700: lipgloss.Color("#3f3f46"),
	800: lipgloss.Color("#27272a"),
	900: lipgloss.Color("#18181b"),
	950: lipgloss.Color("#09090b"),
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
