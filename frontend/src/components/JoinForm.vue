<template>
  <v-col>
    <v-form id="form" v-model="valid" lazy-validation onsubmit="return false;">
      <span class="display-2 display-sm-3 text-center d-block mb-3">Регистрация</span>

      <v-text-field v-model="username" :rules="usernameRules" label="Логин" required name="username"/>

      <v-text-field v-model="email" :rules="emailRules" label="Email" required name="email" type="email"/>

      <v-textarea v-model="about" :rules="aboutRules" label="О себе" name="about" rows="2"/>

      <v-text-field v-model="password" :rules="passwordRules" label="Пароль" required name="password" type="password"/>

      <v-text-field
          v-model="repeatPassword" :rules="repeatPasswordRules" label="Повторите пароль"
          required name="repeatPassword" type="password"/>

      <div class="mb-3 sign-in">
        <router-link to="/login">Уже зарегестрированы?</router-link>
      </div>

      <div class="d-flex justify-center">
        <v-btn x-large color="success" class="mx-4" :disabled="!valid" type="submit" @click="submit">
          Зарегистрироваться
        </v-btn>
      </div>
    </v-form>
  </v-col>
</template>

<script>
  export default {
    name: "JoinForm",
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
      email: '',
      emailRules: [
        v => !!v || "Поле должно быть заполнено"
      ],
      about: '',
      aboutRules: [],
      repeatPassword: '',
      repeatPasswordRules: [
        v => !!v || "Поле должно быть заполнено"
      ]
    }),
    methods: {
      submit: function () {
        if (this.password !== this.repeatPassword) {
          return console.log("Password does not match!")
        }

        this.$store.dispatch('register', {
          username: this.username,
          email: this.email,
          about: this.about,
          password: this.password
        })
            .then(() => this.$router.push('/'))
            .catch(err => console.log(err))
      }
    },
  }
  ;
</script>

<style>
  .v-input__slot {
    margin-bottom: 0 !important;
  }
  .sign-in {
    text-align: left;
  }
  .sign-in a {
    text-decoration: none;
  }
  .sign-in a:hover {
    text-decoration: underline;
  }
</style>