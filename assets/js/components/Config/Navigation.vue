<template>
	<nav>
		<router-link to="/" class="p-2 d-flex align-items-center text-decoration-none">
			<shopicon-regular-arrowback class="me-2"></shopicon-regular-arrowback>
			zurück
		</router-link>
		<router-link class="router-link" to="/config/general">
			<shopicon-regular-home class="me-2"></shopicon-regular-home>
			Allgemein
		</router-link>
		<router-link class="router-link" to="/config/loadpoint">
			<shopicon-regular-cablecharge class="me-2"></shopicon-regular-cablecharge>
			Ladepunkt
		</router-link>
		<router-link class="router-link" to="/config/pv">
			<shopicon-regular-sun class="me-2"></shopicon-regular-sun>
			Photovoltaik
		</router-link>
		<router-link class="router-link" to="/config/battery">
			<shopicon-regular-batterythreequarters
				class="me-2"
			></shopicon-regular-batterythreequarters>
			Hausakku
		</router-link>
		<router-link class="router-link" to="/config/vehicle">
			<shopicon-regular-car3 class="me-2"></shopicon-regular-car3>
			Fahrzeug
		</router-link>
		<router-link class="router-link" to="/config/integration">
			<shopicon-regular-share class="me-2"></shopicon-regular-share>
			Integration
		</router-link>
	</nav>
</template>

<script>
import "@h2d2/shopicons/es/regular/arrowback";
import "@h2d2/shopicons/es/regular/home";
import "@h2d2/shopicons/es/regular/cablecharge";
import "@h2d2/shopicons/es/regular/batterythreequarters";
import "@h2d2/shopicons/es/regular/car3";
import "@h2d2/shopicons/es/regular/share";

export default {
	name: "Navigation",
};
</script>
<style scoped>
.router-link {
	display: flex;
	align-items: flex-end;
	text-decoration: none;
	padding: 0.5rem;
	color: var(--evcc-default-text) !important;
}
.router-link:hover {
	color: var(--evcc-gray) !important;
}
.router-link-active {
	color: var(--bs-primary) !important;
}
.router-link-active:hover {
	color: var(--evcc-accent3) !important;
}
</style>
