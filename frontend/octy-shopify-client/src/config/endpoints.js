const octyAppHost = '' // host or ngrok_tunnel set in ./octy-shopify/pkg/config/config.yaml

module.exports = {
  octyAppHost,
  createCustomerURI: octyAppHost + '/api/customers/createupdate',
  createEventURI: octyAppHost + '/api/events/create',
  createItemURI: octyAppHost + '/api/hooks/items/create',
  updateItemURI: octyAppHost + '/api/hooks/items/update',
  deleteItemURI: octyAppHost + '/api/hooks/items/delete',
  getContentURI: octyAppHost + '/api/content',
  getRecommendationsURI: octyAppHost + '/api/recommendations'
}
