<template>
  <v-container class="my-12">
    <v-row
        class="mx-md-6 px-2" justify="center" no-gutters>
      <v-col :lg="8" :md="8" :sm="10" class="px-2">
        <v-row class="mb-6">
          <v-col :lg="3" :md="4" :sm="12" class="user-avatar">
            <v-avatar color="grey" size="200">
              <v-icon dark size="100">mdi-account-circle</v-icon>
            </v-avatar>
          </v-col>

          <v-col :lg="9" :md="8" class="user-info">
            <h1 class="display-2 text-truncate mb-2">Кирилл Петрунин</h1>
            <div class="title font-weight-regular">
              <span class="d-block mb-2 font-weight-light">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Adipisci alias amet, aperiam asperiores at aut ducimus eveniet explicabo ipsam modi molestiae nemo.</span>
              <span class="d-block mb-2">
                  Дата регистрации: <span class="font-weight-light">25.10.2019</span>
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
                    <div class="recent-post ma-2">
                      <div class="recent-post-header">
                        <span class="title text-capitalize">опа гангнам стайл</span>,
                        <span class="subtitle-1 font-weight-light">25.10.2019 14:22</span>
                      </div>
                      <div class="recent-post-description">
                        <div class="subtitle-1">
                          <span>
                            Предмет: <a href="">Математический Анализ</a>
                          </span>
                        </div>
                        <div class="subtitle-1">
                          <span>
                            Преподаватель: <a href="">Клевчихин Юрий Александрович</a>
                          </span>
                        </div>
                      </div>
                    </div>

                    <v-divider/>

                    <div class="recent-post ma-2">
                      <div class="recent-post-header">
                        <span class="title text-capitalize">опа гангнам стайл</span>,
                        <span class="subtitle-1 font-weight-light">25.10.2019 14:22</span>
                      </div>
                      <div class="recent-post-description">
                        <div class="subtitle-1">
                          <span>
                            Предмет: <a href="">Математический Анализ</a>
                          </span>
                        </div>
                        <div class="subtitle-1">
                          <span>
                            Преподаватель: <a href="">Клевчихин Юрий Александрович</a>
                          </span>
                        </div>
                      </div>
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
  export default {
    name: "UserProfile",
    data() {
      return {
        tab: null
      }
    }
  }
</script>

<style scoped>
  .user-avatar {
    max-width: 200px;
    margin-right: 25px
  }

  .user-info a {
    text-decoration: none;
  }

  .user-info a:hover {
    text-decoration: underline;
  }
</style>