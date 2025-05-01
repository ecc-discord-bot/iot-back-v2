// const auth = new AuthKit('/auth/');
const auth = new PocketBase("/auth/");

async function Init() {
    // ログインしているか確認
    if (pb.authStore.isValid) {
        // ログイン済み
        window.location.href = "./home.html";
    }
}

function OauthLogin(provider) {
    auth.OauthLogin(provider,LoginSuccess);
}

function LoginSuccess() {
    window.location.href = "./home.html";
}

Init();