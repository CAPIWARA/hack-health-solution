<template>
  <harsh-screen classy="schedule-screen">
    <schedule-latest :date="date" />
    <schedule-next :date="date" />
  </harsh-screen>
</template>

<script>
  import ScheduleNext from '@/components/Schedule/ScheduleNext';
  import ScheduleLatest from '@/components/Schedule/ScheduleLatest';

  const DAY = 24 * 60 * 60 * 1000;

  export default {
    components: { ScheduleNext, ScheduleLatest },
    data () {
      return {
        date: new Date(Date.now() - 18 * DAY)
      };
    }
  };
</script>

<style lang="stylus">
  .schedule-screen
    padding: 15px 30px

</style>


