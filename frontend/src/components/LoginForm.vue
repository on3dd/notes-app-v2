<template>
  <v-col>
    <v-form id="form" lazy-validation onsubmit="return false;" v-model="valid">
      <span class="display-2 display-sm-3 text-center d-block mb-3">Авторизация</span>

      <v-text-field :rules="usernameRules" label="Логин" name="username" required v-model="username"/>

      <v-text-field :rules="passwordRules" label="Пароль" name="password" required type="password" v-model="password"/>

      <v-checkbox label="Оставаться в системе" v-model="rememberMe"/>

      <div class="mb-3 sign-up">
        <router-link to="/join">Еще не зарегестрированы?</router-link>
      </div>

      <div class="d-flex justify-center">
        <v-btn :disabled="!valid" @click="submit" class="mx-4" color="success" type="submit" x-large>
          Войти
        </v-btn>
      </div>
    </v-form>
  </v-col>
</template>


<script>
  import axios from "axios";

  export default {
    name: "LoginForm",
    data: () => ({
      valid: false,
      username: '',
      usernameRules: [
        v => !!v || "Поле должно быть заполнено"
      ],
      password: '',
      passwordRules: [
        v => !!v || "Поле должно быть заполнено"
      ],
      rememberMe: false,
    }),
    methods: {
      submit: function () {
        let data = new FormData()
        data.append("name", this.username)
        data.append("password", this.password)
        axios.post("http://localhost:8080/api/v1/login", data)
            .then(response => {
              if (response.status === 200) {
                this.$router.replace({path: `/notes`})
              }
            })
            .catch(err => {
              console.error(err)
            })
      }
    }
  };
</script>

<style>
  .v-input__slot {
    margin-bottom: 0 !important;
  }

  .sign-up {
    text-align: left;
  }

  .sign-up a {
    text-decoration: none;
  }

  .sign-up a:hover {
    text-decoration: underline;
  }
</style>