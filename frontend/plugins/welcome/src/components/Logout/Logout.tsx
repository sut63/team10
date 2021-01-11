import React, { FC } from 'react';
import Avatar from '@material-ui/core/Avatar';
import {
  Link,
  FormControl,
  Button
} from '@material-ui/core';
import CssBaseline from '@material-ui/core/CssBaseline';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';


import { useCookies } from 'react-cookie/cjs';//cookie


const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
    alignItems: 'center',
    marginLeft: 150,

  },
  Typography: {
    margin: theme.spacing(1),

    alignItems: 'center',
    marginLeft: 130,

  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));


const Logout: FC<{}> = () => {
  const classes = useStyles();
  const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  


  const Logout = async () => {

    removeCookie('Name', { path: '/' })
    removeCookie('Password', { path: '/' })
    removeCookie('Log', { path: '/' })
    removeCookie('Status', { path: '/' })
    window.location.reload(false)

  }


  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      
      <div className={classes.paper}>
        
        <form className={classes.form} noValidate>
          <FormControl variant="outlined" className={classes.formControl}>
            <Avatar className={classes.avatar}>
              <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5" className={classes.Typography}>
              Logout
            </Typography>
          </FormControl>
        </form>
        
        <form className={classes.form} noValidate>

         
           &emsp;
            <Link href ="/">
            <FormControl variant="outlined" className={classes.formControl}>
              <Button
                onClick={() => {
                  Logout();
                }}
                variant="contained"
                color="primary"
              >
                Logout
             </Button>
            </FormControl>
            </Link>


        </form>
      </div>
    </Container>
  );
};
export default Logout;