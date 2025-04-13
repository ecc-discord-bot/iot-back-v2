export function appendNameClassForm(confirmTermsButton, termsScrollContainer, agreeList) {
  const formWrapper = document.createElement("div");
  formWrapper.style.marginTop = "1.2rem";
  formWrapper.innerHTML = `
    <div class="input-container">
      <label for="user-name">名前：</label>
      <input type="text" id="user-name" />
    </div>
    <div class="input-container">
      <label for="user-class">クラス：</label>
      <input type="text" id="user-class" />
    </div>
  `;
  agreeList.appendChild(formWrapper);

  const userNameInput = formWrapper.querySelector("#user-name");
  const userClassInput = formWrapper.querySelector("#user-class");

  function validateInputs() {
    const name = userNameInput.value.trim();
    const clazz = userClassInput.value.trim();
    const scrolled = termsScrollContainer.scrollTop + termsScrollContainer.clientHeight >= termsScrollContainer.scrollHeight - 5;
    const isValid = name && clazz && scrolled;

    if (isValid) {
      confirmTermsButton.disabled = false;
      confirmTermsButton.classList.remove("btn-disabled");
      confirmTermsButton.classList.add("btn-blue");
    } else {
      confirmTermsButton.disabled = true;
      confirmTermsButton.classList.add("btn-disabled");
    }
  }

  userNameInput.addEventListener("input", validateInputs);
  userClassInput.addEventListener("input", validateInputs);
  termsScrollContainer.addEventListener("scroll", validateInputs);
}
