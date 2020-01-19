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
      <v-btn color="primary" outlined tile to="/login">
        <v-icon class="mr-1">mdi-account-circle</v-icon>
        <span>Войти</span>
      </v-btn>
    </v-app-bar>

    <v-content class="router-view">
      <v-row class="d-flex justify-center">
        <router-view/>
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
    })
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
</style>