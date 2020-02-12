<template>
  <v-container>
    <v-row
        class="mx-md-6 px-2" justify="center" no-gutters>
      <v-col :lg="10" class="px-2">
        <v-row class="mb-6 d-flex">
          <v-col :lg="3" :md="4" :sm="12" class="user-avatar d-flex justify-center align-center">
            <v-avatar color="grey" size="200">
              <v-icon dark size="100">mdi-account-circle</v-icon>
            </v-avatar>
          </v-col>

          <v-col :lg="9" :md="8" class="user-info">
            <h1 class="display-2 mb-2">{{user.name}}</h1>
            <div class="title font-weight-regular">
              <span class="d-block mb-2 font-weight-light">{{user.about}}</span>
              <span class="d-block mb-2">
                  Дата регистрации: <span class="font-weight-light">25.10.2019, 14:22</span>
                </span>
              <span class="d-block mb-2">
                  Всего работ: <a href="">1488</a>
                </span>
            </div>
          </v-col>
        </v-row>

        <v-row>
          <v-card style="width: 100%">
            <v-tabs center-active background-color="primary" dark v-model="tab">
              <v-tab>Последние работы</v-tab>
              <v-tab>Избранные предметы</v-tab>
            </v-tabs>
            <v-tabs-items style="width: 100%" v-model="tab">
              <v-tab-item>
                <v-card flat v-if="!getUserLastNotes.length">
                  <v-card-text>Данный пользователь еще не загружал работы.</v-card-text>
                </v-card>
                <v-card flat v-else>
                  <v-container>
                    <div :key="index" v-for="(item, index) in getUserLastNotes">
                      <UserProfileRecentNote :note="item"/>
                      <v-divider v-if="index !== getUserLastNotes.length - 1"/>
                    </div>
                  </v-container>
                </v-card>
              </v-tab-item>
              <v-tab-item>
                <v-card v-if="!getUserFavoriteSubjects.length" flat>
                  <v-card-text>Данный пользователь еще не загружал работы.</v-card-text>
                </v-card>
                <v-card v-else flat>
                  <v-container>
                    <div :key="index" v-for="(item, index) in getUserFavoriteSubjects">
                      <UserProfileFavoriteSubject :subject="item"/>
                      <v-divider v-if="index !== getUserFavoriteSubjects.length - 1"/>
                    </div>
                  </v-container>
                </v-card>
              </v-tab-item>
            </v-tabs-items>
          </v-card>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import {mapActions, mapGetters} from "vuex";
  import UserProfileRecentNote from "./UserProfileRecentNote";
  import UserProfileFavoriteSubject from "./UserProfileFavoriteSubject";

  export default {
    name: "UserProfile",
    components: {
      UserProfileRecentNote,
      UserProfileFavoriteSubject
    },
    data() {
      return {
        tab: null
      }
    },
    methods: {
      ...mapActions(["fetchUser", "fetchUserLastNotes", "fetchUserFavoriteSubjects"]),
    },
    computed: {
      ...mapGetters(["getUserLastNotes", "getUserFavoriteSubjects"]),
      user: function () {
        return this.$store.getters.getUser
      },
    },
    async mounted() {
      const id = this.$router.currentRoute.params.id

      await this.fetchUser(id)
      await this.fetchUserLastNotes(id)
      await this.fetchUserFavoriteSubjects(id)

      console.log(this.getUserLastNotes)
      console.log(this.getUserFavoriteSubjects)
    }
  }
</script>

<style>
  .v-slide-group__prev--disabled {
    display: none !important;
  }
</style>

<style scoped>
  @media screen and (min-device-width: 900px) {
    .user-avatar {
      max-width: 200px;
      margin-right: 25px
    }
  }

  @media screen and (max-device-width: 900px){
    .user-info {
      text-align: center
    }
  }
  .user-info a {
    text-decoration: none;
  }

  .user-info a:hover {
    text-decoration: underline;
  }
</style>