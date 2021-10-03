// declare dynamic content sections and their default values
const DCSMap = [
  // Dynamic Content Section Map
  {
    section_id: 'banner-title',
    default_value: 'Welcome to our store!'
  },
  {
    section_id: 'banner-cta',
    default_value: 'Buy some stuff!'
  },
  {
    section_id: 'banner-image',
    default_value:
        "<img srcset='//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_375x.jpg?v=1632134922 375w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_750x.jpg?v=1632134922 750w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_1100x.jpg?v=1632134922 1100w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_1500x.jpg?v=1632134922 1500w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_1780x.jpg?v=1632134922 1780w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_2000x.jpg?v=1632134922 2000w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_3000x.jpg?v=1632134922 3000w,//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_3840x.jpg?v=1632134922 3840w,' sizes='100vw' src='//cdn.shopify.com/s/files/1/0598/2578/2953/files/boots-on-blue_1500x.jpg?v=1632134922' loading='lazy' alt='' width='4096' height='2731.0000000000005'>"
  },
  {
    section_id: 'promo-banner', // NOTE: for this section to work, you must add 'octy-promo-banner.liquid' section to your store.
    default_value: '' // if default value, section will not be rendered
  }
]

module.exports = DCSMap
