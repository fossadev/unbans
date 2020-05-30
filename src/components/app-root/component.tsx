import * as React from "react";
import {ThemeProvider, createMuiTheme} from "@material-ui/core";

interface Props {
  children: React.ReactNode;
}

export const AppRoot: React.FunctionComponent<Props> = (
  props: Props,
) => {
  const theme = createMuiTheme();
  return (
    <ThemeProvider theme={theme}>{props.children}</ThemeProvider>
  );
};
