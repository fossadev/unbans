import * as React from "react";
import {AppRoot} from "components/app-root/component";
import {Switch, Route} from "react-router-dom";

export const Router: React.FunctionComponent = () => {
  return (
    <AppRoot>
      <Switch>
        <Route path="/:channel/dashboard" />
      </Switch>
    </AppRoot>
  );
};
