export function appendNameClassForm(confirmTermsButton, termsScrollContainer, agreeList) {
  const formWrapper = document.createElement("div");
  formWrapper.style.marginTop = "1.2rem";
  formWrapper.innerHTML = `
    <div class="input-container">
      <label for="userName">名前：</label>
      <input type="text" id="userName" />
    </div>
    <div class="input-container">
      <label for="userClass">クラス：</label>
      <input type="text" id="userClass" />
    </div>
  `;
  agreeList.appendChild(formWrapper);

  const userNameInput = formWrapper.querySelector("#userName");
  const userClassInput = formWrapper.querySelector("#userClass");

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
