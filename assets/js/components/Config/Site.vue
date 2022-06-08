<template>
	<form @submit.prevent="save">
		<h2>Site</h2>
		<label for="siteTitle">Titel</label>
		<input id="siteTitle" v-model="form.title" type="text" />
		<button type="submit">Speichern</button>
	</form>
</template>

<script>
import api from "../../api";

export default {
	name: "Site",
	data() {
		return { form: {}, error: null };
	},
	mounted() {
		this.load();
	},
	methods: {
		async load() {
			try {
				this.form = await api.get("config/site").data;
			} catch (e) {
				this.error = this.$t("config.error.loadingFailed");
			}
		},
		async save() {
			try {
				await api.post("config/site", this.form);
			} catch (e) {
				this.error = this.$t("config.error.savingFailed");
			}
		},
	},
};
</script>

<style scoped></style>
