package images

// 颜色串解析(胶水,非算法):"RRGGBB" / "RRGGBB-偏色" / "A|B" 多选 → []kit.Spec。

// simBase sim(0..1)→ 每通道附加容差 round((1-sim)*255)。

// parseAbsMulti "x,y,color,..." → 绝对点(DetectsMultiColors 用)。

// parseRelMulti "anchorColor,dx,dy,color,..." → 锚点候选 + 相对点。
