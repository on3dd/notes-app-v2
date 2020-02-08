<template>
  <a :href="'/notes/'+note.id" class="recent-post pa-2 d-block">
    <div class="recent-post-header">
      <span class="title text-capitalize" style="color: #3e2723">{{note.title}}</span>,
      <span class="subtitle-1 font-weight-light">{{note.posted_at}}</span>
    </div>
    <div class="recent-post-description">
      <div class="subtitle-1">
        <span>
          Предмет: <a href="">{{note.subject_name || "Математический анализ"}}</a>
        </span>
      </div>
      <div class="subtitle-1">
        <span>
          Преподаватель: <a href="">{{note.teacher_name || "Клевчихин Юрий Александрович"}}</a>
        </span>
      </div>
    </div>
  </a>
</template>

<script>
  export default {
    name: "UserProfileRecentNote",
    props: {
      note: {
        required: true
      }
    }
  }
</script>

<style scoped lang="scss">
  .recent-post {
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