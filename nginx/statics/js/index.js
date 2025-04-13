import { setupUI } from "./ui-handlers.js";

// auth
const auth = new AuthKit('/auth/');

function LoginSuccess() {
    EnableMicrosoft();
}

function EnableMicrosoft() {
    // discord を無効か
    loginButton.disabled = true;
    loginButton.classList.add("btn-disabled");

    // microsoft を有効化
    loginButton2.disabled = false;
    loginButton2.classList.remove("btn-disabled");
}

function DisableMicrosoft() {
    EnableMicrosoft();

    // microsoft を無効化
    loginButton2.disabled = true;
    loginButton2.classList.add("btn-disabled");

    // viewTermsButton を有効化
    viewTermsButton.disabled = false;
    viewTermsButton.classList.remove("btn-disabled");
}

async function postStudentInfo(userName, userClass) {
    // 同意ボタンを押したときのイベント
    try {
        await auth.Post("/app/terms/accept", {
            "UserName": userName,
            "UserClass": userClass
        });
    } catch (error) {
        console.error(error);
    }
}

async function Init() {
    // ログインしているか確認
    if (await auth.GetInfo() != null) {
        // ログイン済み
        EnableMicrosoft();
    }

    try {
        // リンク状態を取得
        const req = await auth.Get("/app/link/status", {});

        // ボタン２を有効化

        // 同意済みか
        if (req["IsAgreed"]) {
            // 同意済みの場合
        } else {
            // 同意していない場合
            DisableMicrosoft();
        }
    } catch (error) {
        console.error(error);
    }
}

function OauthLogin(provider) {
    auth.OauthLogin(provider, LoginSuccess);
}


// ボタンを取得
const loginButton = document.getElementById("loginButton1");
const loginButton2 = document.getElementById("loginButton2");
const viewTermsButton = document.getElementById("view-terms-button");

document.addEventListener("DOMContentLoaded", () => {
    // UI
    setupUI(postStudentInfo);

    // Discordログインボタンを押したとき
    loginButton.addEventListener("click", function () {
        OauthLogin("discord");
    });

    // Microsoftログインボタンを押したとき
    loginButton2.addEventListener("click", async function () {
        // キャッシュ状態のトークンを削除
        window.sessionStorage.removeItem("actoken");
        window.sessionStorage.removeItem("actime");

        // セッションを張る
        const req = await auth.Post("/app/startlink", {}, {});

        // リダイレクト
        window.location.href = "/app/aclink";
    });

    Init();
});
