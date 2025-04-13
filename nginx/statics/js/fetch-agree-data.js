import { renderBlocks } from "./render-block.js";
import { appendNameClassForm } from "./form-validation.js";

export async function fetchAgreeData(agreeList, confirmTermsButton, termsScrollContainer) {
  agreeList.innerHTML = '<div class="loading-overlay visible"><div class="spinner"></div><div>規約を読み込んでいます...</div></div>';
  try {
    const res = await fetch("https://localhost:8370/statics/terms.json", {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    });
    const data = await res.json();

    await new Promise((resolve) => setTimeout(resolve, 500));
    agreeList.innerHTML = "";
    data.forEach((page) => {
      const title = page.properties.number?.title?.[0]?.plain_text || "（無題）";
      const summary = page.properties.title?.rich_text?.[0]?.plain_text || "";
      const contentHTML = renderBlocks(page.content || []);
      const item = document.createElement("div");
      item.style.marginBottom = "1rem";
      item.style.paddingBottom = "0.4rem";
      item.style.borderBottom = "1px solid #d1d5db";
      item.innerHTML = `
        <strong style="font-size: 1.2rem;">${title}：${summary}</strong>
        ${contentHTML ? `<div style="font-size: 0.8rem;">${contentHTML}</div>` : ""}
      `;
      agreeList.appendChild(item);
    });

    appendNameClassForm(confirmTermsButton, termsScrollContainer, agreeList);
  } catch (error) {
    agreeList.innerHTML = '<p style="font-size: 0.8rem; color: red;">規約の読み込みに失敗しました。</p>';
    confirmTermsButton.disabled = true;
  }
}
