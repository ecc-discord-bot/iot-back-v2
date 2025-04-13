const auth = new AuthKit('/auth/'); 

async function Init() {    
    try {
        // 情報を取得
        const info = await auth.GetInfo();

        if (info == null) {
            // ログインにリダイレクト
            window.location.href = './login.html';
            return;
        }

        console.log(info);

        // hello world 検証
        const hello = await auth.Get('/app/hello',{});

        console.log(hello);
    } catch (error) {
        console.error(error);
        // ログインにリダイレクト
        window.location.href = './login.html';
    }
}

Init();