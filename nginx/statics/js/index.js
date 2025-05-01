import { setupUI } from "./ui-handlers.js";

// auth
const auth = new AuthKit('/auth/');

function LoginSuccess() {
    EnableMicrosoft();

    // リロード
    window.location.reload();
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
    console.log(userName, userClass);

    // 同意ボタンを押したときのイベント
    try {
        await auth.Post("/app/terms/accept", {
            "Content-Type": "application/json"
        }, JSON.stringify({
            "UserName": userName,
            "UserClass": userClass
        }));
    } catch (error) {
        console.error(error);
    }
}

async function Init() {
    // UI 初期化
    await InitUI(false);

    // ログインしているか確認
    if (await auth.IsValid()) {
        // ログイン済み
        EnableMicrosoft();
    } else {
        // ログインしていない
        return;
    }

    try {
        // リンク状態を取得
        const req = await auth.Get("/app/link/status", {});

        // ボタン２を有効化

        // 同意済みか
        if (req["IsAgreed"]) {
            // 同意済みの場合
            await InitUI(true);
            
            // microsoft を無効化
            DisableMicrosoft();

            return;
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

async function InitUI(isAgreed) {
    try {
        // リンクを取得
        const discordLink = await auth.Get("/app/invite", {});
            
        // UI
        setupUI(postStudentInfo,isAgreed,discordLink["url"]);
    } catch (error) {
        console.error(error);
        setupUI(postStudentInfo,isAgreed,"");
    }


    // Discordログインボタンを押したとき
    loginButton.addEventListener("click", function () {
        OauthLogin("discord");
    });

    // Microsoftログインボタンを押したとき
    loginButton2.addEventListener("click", async function () {
        try {
            // キャッシュ状態のトークンを削除
            window.sessionStorage.removeItem("actoken");
            window.sessionStorage.removeItem("actime");

            // セッションを張る
            const req = await auth.Post("/app/startlink", {}, {});

            // リダイレクト
            window.location.href = "/app/aclink";
        } catch (error) {
            console.error(error);
        }
    });
};

Init();