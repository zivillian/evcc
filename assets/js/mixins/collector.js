export default {
  methods: {
    // collect all target component properties from current instance
    collectProps: function (component) {
      let data = this.collectComponentProps(component);
      for (var c in component.components) {
        console.log(c);
        console.log(component.components[c]);
        Object.assign(data, this.collectProps(component.components[c]));
      }
      console.log(data);
      return data;
    },
    collectComponentProps: function (component) {
      let data = {};
      for (var p in component.props) {
        if (p in this) {
          data[p] = this[p];
        }
      }
      return data;
    },
  },
};
