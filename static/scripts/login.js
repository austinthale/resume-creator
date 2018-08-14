new Vue({
    el: '#login',
    data:{
        account:{
            username: "",
            password: "",
        },
        isLoggedIn: false,
    },
    methods: {
        loginButton: function () {
            validateLogin();
            //this.isLoggedIn = true;
        },
    }
}