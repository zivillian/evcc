<template>
	<form @submit.prevent="save">
		<label for="siteTitle">Site Title</label>
		<input id="siteTitle" v-model="form.title" type="text" />
		<button type="submit">Speichern</button>
	</form>
</template>

<script>
import api from "../../api";

export default {
	name: "Site",
	data() {
		return { form: { title: null }, error: null };
	},
	mounted() {
		this.load();
	},
	methods: {
		async load() {
			try {
				const response = await api.get("config/site/title");
				this.form.title = response.data;
			} catch (e) {
				this.error = this.$t("config.error.loadingFailed");
			}
		},
		async save() {
			try {
				await api.put("config/site/title", this.form.title);
			} catch (e) {
				this.error = this.$t("config.error.savingFailed");
			}
		},
	},
};
</script>

<style scoped></style>
