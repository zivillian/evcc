<template>
	<div class="container px-4">
		<header class="py-3 d-flex align-items-center justify-content-between">
			<h1 class="my-0">Configuration</h1>
			<router-link v-slot="{ navigate }" to="/" custom>
				<button type="button" class="btn btn-sm btn-outline-secondary" @click="navigate">
					<shopicon-regular-arrowback></shopicon-regular-arrowback>
				</button>
			</router-link>
		</header>

		<main>
			<Site />
		</main>
	</div>
</template>

<script>
import "@h2d2/shopicons/es/regular/arrowback";

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
