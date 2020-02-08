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
            <h1 class="display-2 text-truncate mb-2">{{user.name}}</h1>
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
            <v-tabs background-color="primary" dark v-model="tab">
              <v-tab>Последние работы</v-tab>
              <v-tab>Избранные предметы</v-tab>
            </v-tabs>
            <v-tabs-items style="width: 100%" v-model="tab">
              <v-tab-item>
                <v-card flat>
                  <v-container>
                    <div :key="index" v-for="(item, index) in getUserLastNotes">
                      <v-divider v-if="index % 2 === 1"/>
                      <UserProfileRecentNote :note="item"/>
                      <v-divider v-if="index % 2 === 1"/>
                    </div>
                  </v-container>
                </v-card>
              </v-tab-item>
              <v-tab-item>
                <v-card flat>
                  <v-card-text>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aliquam, sed.</v-card-text>
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

  export default {
    name: "UserProfile",
    components: {UserProfileRecentNote},
    data() {
      return {
        tab: null
      }
    },
    methods: {
      ...mapActions(["fetchUser", "fetchUserLastNotes"]),
    },
    computed: {
      ...mapGetters(["getUserLastNotes"]),
      user: function () {
        return this.$store.getters.getUser
      },
    },
    async mounted() {
      const id = this.$router.currentRoute.params.id
      await this.fetchUser(id)
      await this.fetchUserLastNotes(id)
      console.log(this.getUserLastNotes)
    }
  }
</script>

<style scoped>
  @media screen and (min-device-width: 900px) {
    .user-avatar {
      max-width: 200px;
      margin-right: 25px
    }
  }

  .user-info a {
    text-decoration: none;
  }

  .user-info a:hover {
    text-decoration: underline;
  }
</style>