<template>
    <v-footer
        color="primary lighten-1"
        padless
    >
      <v-row
          justify="center"
          no-gutters
      >
        <v-btn
            :key="link"
            class="my-2"
            color="white"
            rounded
            text
            v-for="link in links"
        >
          {{ link }}
        </v-btn>
        <v-col
            class="primary lighten-2 py-4 text-center white--text"
            cols="12"
        >
          {{ new Date().getFullYear() }} — <strong>Vuetify</strong>
        </v-col>
      </v-row>
    </v-footer>
</template>

<script>
  export default {
    name: "Footer",
    data: () => ({
      links: [
        'Home',
        'About Us',
        'Team',
        'Services',
        'Blog',
        'Contact Us',
      ],
    })
  }
</script>