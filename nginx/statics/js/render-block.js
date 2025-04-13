export function renderBlocks(blocks) {
  if (!blocks || blocks.length === 0) return "";
  return blocks
    .map((block) => {
      const richText = block[block.type]?.rich_text || [];
      const text = richText.map((t) => t.text?.content).join("");
      let childrenHTML = block.children?.length ? `<ul style="padding: 0;">${renderBlocks(block.children)}</ul>` : "";
      if (block.type === "bulleted_list_item") {
        return `<ul><li>${text}${childrenHTML}</li></ul>`;
      } else if (block.type === "paragraph") {
        return `<p>${text}${childrenHTML}</p>`;
      } else {
        return `<div>${text}${childrenHTML}</div>`;
      }
    })
    .join("");
}
