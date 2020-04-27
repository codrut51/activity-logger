import { createMuiTheme } from '@material-ui/core/styles';

// A custom theme for this app
const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#63adf2',
    },
    secondary: {
      main: '#a7cced',
    },
    error: {
      main: '#af2bbf',
    },
    background: {
      default: '#fff',
    },
  },
});

export default theme;