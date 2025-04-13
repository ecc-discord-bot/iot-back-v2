import { fetchAgreeData } from "./fetch-agree-data.js";

export function setupUI(postStudentInfo,isAgreed) {
  const card = document.getElementById("card");
  const perspectiveContainer = document.getElementById("perspective-container");
  const viewTermsButton = document.getElementById("view-terms-button");
  const confirmTermsButton = document.getElementById("confirm-terms-button");
  const termsScrollContainer = document.getElementById("terms-scroll-container");
  const agreeList = document.getElementById("agree-list");
  const loadingOverlay = document.getElementById("loadingOverlay");

  viewTermsButton.addEventListener("click", () => {
    confirmTermsButton.disabled = true;
    card.classList.add("flipped");
    perspectiveContainer.classList.add("flipped-container");
    fetchAgreeData(agreeList, confirmTermsButton, termsScrollContainer);
  });

  confirmTermsButton.addEventListener("click", async () => {
    const userName = document.getElementById("userName")?.value?.trim();
    const userClass = document.getElementById("userClass")?.value?.trim();
    if (!userName || !userClass) return alert("名前とクラスを入力してください");

    try {
      loadingOverlay.classList.add("visible");
      await postStudentInfo(userName, userClass);
      card.classList.remove("flipped");
      perspectiveContainer.classList.remove("flipped-container");

      const frontCard = card.querySelector(".card-front");
      frontCard.innerHTML = `
        <div class="text-center">
          <h2 class="text-green">ログイン成功！</h2>
          <p class="text-gray">ようこそ！</p>
        </div>`;
    } catch (e) {
      alert("送信に失敗しました");
    } finally {
      loadingOverlay.classList.remove("visible");
    }
  });

  if (isAgreed) {
    loadingOverlay.classList.add("visible");
    card.classList.remove("flipped");
    perspectiveContainer.classList.remove("flipped-container");

    const frontCard = card.querySelector(".card-front");
    frontCard.innerHTML = `
      <div class="text-center">
        <h2 class="text-green">ログイン成功！</h2>
        <p class="text-gray">ようこそ！</p>
      </div>`;
    
    loadingOverlay.classList.remove("visible");
  }
}
