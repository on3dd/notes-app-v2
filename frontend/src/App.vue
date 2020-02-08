<template>
  <v-app>
    <v-navigation-drawer
        app
        v-model="drawer"
    >
      <v-list dense>
        <v-list-item :key="index" :to="item.path" v-for="(item, index) in items">
          <v-list-item-action>
            <v-icon>{{item.icon}}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title style="font-size: 15px">{{item.text}}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider/>

        <v-list-item to="/upload">
          <v-list-item-action>
            <v-icon>mdi-upload</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title style="font-size: 15px">Загрузка</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>
      <v-toolbar-title>Конспекты</v-toolbar-title>

      <v-spacer/>

      <div v-if="!isLoggedIn">
        <v-btn color="primary" outlined tile to="/login">
          <v-icon class="mr-1">mdi-account-circle</v-icon>
          <span>Войти</span>
        </v-btn>
      </div>

      <div v-else>
        <v-btn :to="`/users/${user_id}`" color="primary" outlined tile class="mr-2" title="Перейти в профиль">
          <v-icon class="mr-1">mdi-account-circle</v-icon>
          <span>{{user_name}}</span>
        </v-btn>

        <v-btn @click="logout" tile icon color="primary" class="logout-btn" title="Выход">
          <v-icon>mdi-logout</v-icon>
        </v-btn>
      </div>

    </v-app-bar>

    <v-content class="router-view">
      <v-row class="d-flex justify-center">
        <router-view style="margin: 20px 0"/>
      </v-row>
    </v-content>

    <v-content class="pt-0">
      <Footer/>
    </v-content>
  </v-app>
</template>

<script>
  import Footer from "./components/Footer";

  export default {
    name: "App",
    components: {
      Footer
    },
    data: () => ({
      drawer: true,
      items: [
        {path: "/", icon: "mdi-home", text: "Главная"},
        {path: "/notes", icon: "mdi-folder-open", text: "Работы"},
        {path: "/users", icon: "mdi-account-group", text: "Пользователи"}
      ]
    }),
    methods: {
      logout: function () {
        this.$store.dispatch('logout')
            .then(() => {
              this.$router.push('/login')
            })
      }
    },
    computed: {
      isLoggedIn: function () {
        return this.$store.getters.isLoggedIn
      },
      user: function () {
        return this.$store.getters.getCurrentUser
      },
      user_name: function () {
        const user_name = localStorage.getItem('user_name')
        return user_name === null ? 'Аноним' : user_name
      },
      user_id: function () {
        const user_id = localStorage.getItem('user_id')
        return user_id === null ? '0' : user_id
      },
    },
    created() {
      this.$http.interceptors.response.use(undefined, function (err) {
        return new Promise(function (resolve, reject) {
          if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
            this.$store.dispatch("logout")
          }
          throw err;
        });
      });
    },
    watch: {
      '$route'(to, from) {
        document.title = to.meta.title || 'Конспекты'
      }
    },
  };
</script>

<style>
  body {
    overflow-x: hidden;
  }

  .router-view {
    min-height: 100vh;
  }

  .v-list-item--active .v-icon {
    color: #1976d2 !important;
  }

  .logout-btn {
    max-height: 36px;
    max-width: 36px;
  }
</style>