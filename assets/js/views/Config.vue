<template>
	<div>
		<header class="container px-4 d-flex justify-content-between">
			<h1>evcc Configuration</h1>
			<router-link to="/">zurück</router-link>
		</header>

		<main>
			<Site />
		</main>
	</div>
</template>

<script>
import Site from "../components/Config/Site.vue";

export default {
	name: "Config",
	components: { Site },
	props: {
		notifications: Array,
	},
};
</script>
<style scoped></style>
