export function launchConfetti(duration = 3000, count = 30) {
  const container = document.getElementById("confetti-container");

  for (let i = 0; i < count; i++) {
    const confetti = document.createElement("div");
    confetti.classList.add("confetti");

    const size = Math.random() * 10 + 10;
    confetti.style.width = `${size}px`;
    confetti.style.height = `${size}px`;

    confetti.style.backgroundImage = "url(/statics/assets/sakura.webp)";
    confetti.style.backgroundSize = "cover";
    confetti.style.fontSize = `${size}px`;
    confetti.style.textAlign = "center";

    confetti.style.left = `${Math.random() * 100}%`;
    confetti.style.animationDuration = `${Math.random() * 2 + 3}s`;
    confetti.style.animationDelay = `${Math.random() * 1}s`;

    container.appendChild(confetti);

    // アニメーション後に削除
    setTimeout(() => container.removeChild(confetti), duration);
  }
}
