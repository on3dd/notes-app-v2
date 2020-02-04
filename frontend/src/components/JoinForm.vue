<template>
  <v-col>
    <ValidationObserver ref="obs" v-slot="{ invalid, validated, passes, validate }">
      <v-form id="form" lazy-validation onsubmit="return false;">
        <span class="display-2 display-sm-3 text-center d-block mb-3">Регистрация</span>

        <VTextFieldWithValidation :counter="20" label="Имя пользователя" rules="required|max:20" v-model="username"/>

        <VTextFieldWithValidation label="Электронная почта" rules="required|email" v-model="email"/>

        <v-textarea label="О себе" name="about" rows="2" v-model="about"/>

        <ValidationObserver>
          <VTextFieldWithValidation label="Пароль" rules="required|min:8|password:@Повторите пароль" type="password"
                                    v-model="password"/>

          <VTextFieldWithValidation label="Повторите пароль" name="confirm" rules="required|min:8" type="password"
                                    v-model="repeatPassword"/>
        </ValidationObserver>

        <div class="mb-3 sign-in">
          <router-link to="/login">Уже зарегестрированы?</router-link>
        </div>

        <div class="d-flex justify-center">
          <v-btn :disabled="invalid" @click="submit" class="mx-4" color="success" type="submit" x-large>
            Зарегистрироваться
          </v-btn>
        </div>
      </v-form>
    </ValidationObserver>
  </v-col>
</template>

<script>
  import {ValidationObserver} from "vee-validate";
  import VTextFieldWithValidation from './inputs/VTextFieldWithValidation';

  export default {
    name: "JoinForm",
    components: {
      ValidationObserver,
      // ValidationProvider,
      VTextFieldWithValidation
    },
    data: () => ({
      valid: false,

      username: '',
      email: '',
      emailRules: [
        v => !!v || "Поле должно быть заполнено",
        v => /.+@.+/.test(v) || 'Некорректный адрес',
      ],

      about: '',

      password: '',
      passwordRules: [
        v => !!v || "Поле должно быть заполнено"
      ],

      repeatPassword: '',
      repeatPasswordRules: [
        v => !!v || "Поле должно быть заполнено",
      ]
    }),
    methods: {
      submit: function () {
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
  };
</script>

<style>
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