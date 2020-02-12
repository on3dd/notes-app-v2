<template>
  <a :href="'/subjects/'+subject.id" class="favorite-subject pa-2 d-block">
    <div class="favorite-subject-header">
      <span class="title text-capitalize" style="color: #3e2723">{{subject.name}}</span>
    </div>
    <div class="favorite-subject-description">
      <div class="subtitle-1">
        <span>
          Категория: <a href="">{{subject.category_name || "Математика"}}</a>
        </span>
      </div>
      <div class="subtitle-1">
        <span>
          Всего работ: <a href="">{{subject.posts_count}}</a>
        </span>
      </div>
    </div>
  </a>
</template>

<script>
  export default {
    name: "UserProfileFavoriteSubject",
    props: {
      subject: {
        required: true
      }
    }
  }
</script>

<style scoped lang="scss">
  .favorite-subject {
    color: #3e2723;
    text-decoration: none;

    &:hover {
      background-color: darken(#fafafa, 2.5%);
    }

    a {
      text-decoration: none;
      &:hover {
        text-decoration: underline;

      }
    }
  }
</style>