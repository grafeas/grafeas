
# Layer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**directive** | [**DirectiveEnum**](#DirectiveEnum) | The recovered Dockerfile directive used to construct this layer. |  [optional]
**arguments** | **String** | The recovered arguments to the Dockerfile directive. |  [optional]


<a name="DirectiveEnum"></a>
## Enum: DirectiveEnum
Name | Value
---- | -----
UNKNOWN_DIRECTIVE | &quot;UNKNOWN_DIRECTIVE&quot;
MAINTAINER | &quot;MAINTAINER&quot;
RUN | &quot;RUN&quot;
CMD | &quot;CMD&quot;
LABEL | &quot;LABEL&quot;
EXPOSE | &quot;EXPOSE&quot;
ENV | &quot;ENV&quot;
ADD | &quot;ADD&quot;
COPY | &quot;COPY&quot;
ENTRYPOINT | &quot;ENTRYPOINT&quot;
VOLUME | &quot;VOLUME&quot;
USER | &quot;USER&quot;
WORKDIR | &quot;WORKDIR&quot;
ARG | &quot;ARG&quot;
ONBUILD | &quot;ONBUILD&quot;
STOPSIGNAL | &quot;STOPSIGNAL&quot;
HEALTHCHECK | &quot;HEALTHCHECK&quot;
SHELL | &quot;SHELL&quot;



