const toggleButton = document.getElementById("toggle-steps-button");
const closeButton = document.getElementById("close-steps-button");
const stepsCurtain = document.getElementById("steps-curtain");

// デフォルトで表示しておく
stepsCurtain.classList.add("visible");

toggleButton.addEventListener("click", () => {
  stepsCurtain.classList.toggle("visible");
  if (stepsCurtain.classList.contains("visible")) {
    toggleButton.textContent = "手順を隠す";
  } else {
    toggleButton.textContent = "手順を確認する";
  }
});

closeButton.addEventListener("click", () => {
  stepsCurtain.classList.remove("visible");
  toggleButton.textContent = "手順を確認する";
});

stepsCurtain.addEventListener("click", (event) => {
  if (event.target === stepsCurtain) {
    stepsCurtain.classList.remove("visible");
    toggleButton.textContent = "手順を確認する";
  }
});
