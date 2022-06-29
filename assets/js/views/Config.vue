<template>
	<div class="container px-4">
		<header class="py-3">
			<h1 class="my-1">Configuration</h1>
		</header>

		<div class="row">
			<Navigation class="col-4" />
			<main class="col-8">
				<router-view />
			</main>
		</div>
	</div>
</template>

<script>
import Navigation from "../components/Config/Navigation.vue";

export default {
	name: "Config",
	components: { Navigation },
	props: {
		notifications: Array,
	},
};
</script>
<style scoped></style>
