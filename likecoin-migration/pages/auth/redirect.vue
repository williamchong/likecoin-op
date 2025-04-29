<script>
export default {
  fetch({ redirect, query }) {
    const { code, method, ...others } = query;
    const queries = new URLSearchParams({
      code,
      method,
      ...others,
    });
    redirect(`/?${queries.toString()}`);
  },
};
</script>
