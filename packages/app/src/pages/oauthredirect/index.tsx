import { navigate, PageProps } from 'gatsby';
import { parse } from 'query-string';
import React from 'react';
import { useRecoilState } from 'recoil';
import Request from '$components/request';
import useTheme from '$hooks/theme';
import { KEY, get } from '$storage/index';
import { oneState as endpointOneState } from '$store/selectors/endpoint';
import { EndpointID } from '$types/index';
import {
  Document,
  RequestPayloadParameter,
  RequestPayloadRequestBody,
} from '$types/oas';
import { promiseErrorHandler } from '$utils/index';
import {
  constructDefaultValues,
  constructFakeDocument,
  constructRequestInfo,
  constructRequestInit,
  getRequest,
} from '$utils/oas/index';

type Props = PageProps;
const OAuthRedirectPage: React.FC<Props> = ({ location }) => {
  useTheme();
  const queries = parse(location.search);
  const endpointId = get<EndpointID>(KEY.OAUTH_ENDPOINT_ID);
  const [endpoint] = useRecoilState(endpointOneState({ id: endpointId }));

  if (!endpoint) {
    throw new Error('Endoint Not Found.');
  }

  const authConfig = endpoint.authConfigs.find(function (authConfig) {
    return authConfig.type === 'oauthcallback';
  });

  if (!authConfig) {
    throw new Error('AuthConfig of type oauthcallback not found.');
  }

  const { pathObject } = authConfig;
  const document: Document = constructFakeDocument({ paths: pathObject });
  const request = getRequest(document);

  if (!request) {
    throw new Error('TODO');
  }

  const defaultValues = constructDefaultValues(request, queries, queries);

  const handleSubmit = async function ({
    requestPayloadParameters,
    requestPayloadRequestBody,
  }: {
    requestPayloadParameters?: RequestPayloadParameter[];
    requestPayloadRequestBody?: RequestPayloadRequestBody;
  } = {}) {
    const requestInfo: RequestInfo = constructRequestInfo(
      endpoint,
      document,
      request,
      requestPayloadParameters
    );
    const requestInit: RequestInit = constructRequestInit(
      request,
      requestPayloadParameters,
      requestPayloadRequestBody
    );
    const [response, responseError] = await promiseErrorHandler(
      fetch(requestInfo, requestInit)
    );
    if (!!responseError) {
      // TODO
      return;
    }
    if (!response.ok) {
      // TODO
      return;
    }
    navigate(`/endpoints/${endpoint.id}`);
  };

  return (
    <div id="page-oauthredirect">
      <p>Processing OAuth redirection...</p>
      <p>{`https://localhost:8000/oauthredirect`}</p>
      <Request
        request={request}
        defaultValues={defaultValues}
        onSubmit={handleSubmit}
      />
    </div>
  );
};
export default OAuthRedirectPage;
