module.exports = {
  chainWebpack (configuration) {
    configuration.module.rules.delete('graphql');

    configuration.module.rule('graphql')
      .test(/\.graphql$/)
      .use('graphql-raw-loader')
        .loader('graphql-raw-loader');
  }
};
