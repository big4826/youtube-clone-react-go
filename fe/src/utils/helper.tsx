import qs from "qs";

type ParamType = { [p: string]: unknown } | undefined;
type Option<RouteParam extends ParamType, QueryParam extends ParamType> = {
  routeParam?: RouteParam;
  queryParam?: QueryParam;
};

export const customGeneratePath =
  <
    RouteParam extends ParamType = undefined,
    QueryParam extends ParamType = undefined
  >(
    url: string
  ) =>
  (option?: Option<RouteParam, QueryParam>) => {
    const { routeParam, queryParam } = option || {};
    let newQueryParam = "";
    if (queryParam) {
      newQueryParam = `?${qs.stringify(queryParam)}`;
    }

    let newUrl = url;
    if (routeParam) {
      const urls = url.split("/");
      newUrl = urls
        .map((u) => {
          if (/:.+/.test(u)) {
            return routeParam[u.slice(1)];
          }
          return u;
        })
        .join("/");
    }

    return `${newUrl}${newQueryParam}`;
  };
