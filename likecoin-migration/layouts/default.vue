<template>
  <ErrorBoundary>
    <div :class="['flex', 'flex-col', 'min-h-screen']">
      <div :class="['w-full']">
        <PageHeader />
      </div>
      <Nuxt />
    </div>
  </ErrorBoundary>
</template>
