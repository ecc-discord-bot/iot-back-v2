import { fetchAgreeData } from "./fetch-agree-data.js";
import { launchConfetti } from "./launch-confetti.js";

export function setupUI(postStudentInfo, isAgreed, ServerUrl) {
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
    const userName = document.getElementById("user-name")?.value?.trim();
    const userClass = document.getElementById("user-class")?.value?.trim();
    if (!userName || !userClass) return alert("名前とクラスを入力してください");

    try {
      loadingOverlay.classList.add("visible");
      await postStudentInfo(userName, userClass);
      setInterval(() => launchConfetti(8000, 10), 500);
      card.classList.remove("flipped");
      perspectiveContainer.classList.remove("flipped-container");

      const frontCard = card.querySelector(".card-front");
      frontCard.innerHTML = `
        <div class="text-center">
          <h2 class="text-gray">ようこそIoT部！</h2>
          <a href="${ServerUrl}" class="btn btn-blue">IoT部へ</a>
        </div>`;
    } catch (e) {
      alert("送信に失敗しました");
    } finally {
      loadingOverlay.classList.remove("visible");
    }
  });

  if (isAgreed) {
    loadingOverlay.classList.add("visible");
    setInterval(() => launchConfetti(8000, 10), 500);
    card.classList.remove("flipped");
    perspectiveContainer.classList.remove("flipped-container");

    const frontCard = card.querySelector(".card-front");
    frontCard.innerHTML = `
      <div class="text-center">
        <h2 class="text-gray">ようこそIoT部！</h2>
        <a href="${ServerUrl}" class="btn btn-blue">IoT部のDiscordへ参加</a>
      </div>`;
    loadingOverlay.classList.remove("visible");
  }
}
