<<<<<<< HEAD
import forOwn from 'mout/object/forOwn';
import find from 'mout/object/find';
import keys from 'mout/object/keys';
=======
import forEach from 'mout/array/forEach';
import sortBy from 'mout/array/sortBy';
import isNumber from 'mout/lang/isNumber';
import forOwn from 'mout/object/forOwn';
>>>>>>> 3b1db76b95f960f0f0f8bf27167464e0493c22a0
import ObjectAssign from 'object-assign';
import shortid from 'shortid';
import storage from 'store';
import { constants as states } from '../states';

export default {
  /**
   * 1件のエンドポイントを追加します。
   * @param {riotx.Context} context
   * @param {String} endpointKey
   * @param {Object} endpoint
   * @return {Array}
   */
  add: (context, endpointKey, endpoint) => {
    context.state.endpoints[endpointKey] = endpoint;
    storage.set('endpoints', context.state.endpoints);
    return [states.ENDPOINTS];
  },

  /**
   * 指定されたエンドポイントを削除します。
   * @param {riotx.Context} context
   * @param {String} endpointKey
   * @return {Array}
   */
  remove: (context, endpointKey) => {
    delete context.state.endpoints[endpointKey];
    storage.set('endpoints', context.state.endpoints);
    return [states.ENDPOINTS];
  },

  /**
   * 全てのエンドポイントを削除します。
   * @param {riotx.Context} context
   * @return {Array}
   */
  removeAll: context => {
    context.state.endpoints = {};
    storage.set('endpoints', context.state.endpoints);
    return [states.ENDPOINTS];
  },

  /**
   * 指定されたエンドポイントを更新します。
   * @param {riotx.Context} context
   * @param {String} endpointKey
   * @param {Object} endpoint
   * @return {Array}
   */
  update: (context, endpointKey, endpoint) => {
    if (!endpoint) {
      context.state.endpoints[endpointKey] = null;
    } else {
      context.state.endpoints[endpointKey] = ObjectAssign({}, context.state.endpoints[endpointKey], endpoint);
    }
    storage.set('endpoints', context.state.endpoints);
    return [states.ENDPOINTS];
  },

  /**
   * 指定されたエンドポイントのtokenを更新します。
   * @param {riotx.Context} context
   * @param {String} endpointKey
   * @param {String|null} token
   * @return {Array}
   */
  updateToken: (context, endpointKey, token) => {
    if (!!context.state.endpoints[endpointKey]) {
      context.state.endpoints[endpointKey].token = token;
    }
    storage.set('endpoints', context.state.endpoints);
    return [states.ENDPOINTS];
  },

  /**
   * 新エンドポイント群をmergeします。
   * @param {riotx.Context} context
   * @param {Object} endpoints
   * @return {Array}
   */
  mergeAll: (context, endpoints) => {
    const modifiedEndpoint = ObjectAssign({}, context.state.endpoints);
    let newEndpoints = {};

    forOwn(endpoints, (endpoint) => {
      let duplicatedEndpoint = find(modifiedEndpoint, val => {
        return endpoint.url === val.url;
      });

      if (!duplicatedEndpoint) {
        const key = shortid.generate();
        modifiedEndpoint[key] = endpoint;
      } else {
        const searchKeys = keys(modifiedEndpoint);
        const resultKey = searchKeys.filter((key) => {
          return modifiedEndpoint[key] === duplicatedEndpoint;
        });
        modifiedEndpoint[resultKey] = endpoint;
      }
    });

    context.state.endpoints = modifiedEndpoint;
    storage.set('endpoints', modifiedEndpoint);
    return [states.ENDPOINTS];
  },

  /**
   * エンドポイント群のorder値を整理します。
   * order値が存在しない場合は後方に配置されます。
   * @param {riotx.Context} context
   * @return {Array}
   */
  tidyUpOrder: context => {
    const newEndpoints = ObjectAssign(context.state.endpoints);
    // どのorder値よりも大きいであろう適当な値。
    const bigNumber = 9999;
    let ordered = [];
    forOwn(newEndpoints, (endpoint, key) => {
      ordered.push({
        key,
        order: (isNumber(endpoint.order) ? endpoint.order : bigNumber)
      });
    });
    ordered = sortBy(ordered, obj => {
      return obj.order;
    });
    forEach(ordered, (obj, order) => {
      newEndpoints[obj.key].order = order;
    });
    context.state.endpoints = newEndpoints;
    storage.set('endpoints', newEndpoints);
    return [states.ENDPOINTS];
  },

  /**
   * 指定されたエンドポイントのorder値を変更します。
   * 他エンドポイントのorder値もインクリメントされます。
   * @param {riotx.Context} context
   * @param {String} endpointKey
   * @param {Number} newOrder
   * @return {Array}
   */
  changeOrder: (context, endpointKey, newOrder) => {
    const newEndpoints = ObjectAssign(context.state.endpoints);
    // x番目とx+1番目の中間に配置するために0.5をマイナスしている。
    newEndpoints[endpointKey].order = newOrder - 0.5;
    let ordered = [];
    forOwn(newEndpoints, (endpoint, key) => {
      ordered.push({
        key,
        order: endpoint.order
      });
    });
    ordered = sortBy(ordered, obj => {
      return obj.order;
    });
    forEach(ordered, (obj, order) => {
      newEndpoints[obj.key].order = order;
    });
    context.state.endpoints = newEndpoints;
    storage.set('endpoints', newEndpoints);
    return [states.ENDPOINTS];
  }
};
