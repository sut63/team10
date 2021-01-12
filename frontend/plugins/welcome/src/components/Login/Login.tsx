import React, { FC, useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import {
  FormControl,
  Link,
  Button
} from '@material-ui/core';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi } from '../../api/apis';
import { Alert ,AlertTitle } from '@material-ui/lab';


import { useCookies } from 'react-cookie/cjs';//cookie

import { EntUser } from '../../api/models/EntUser';
import { Link as RouterLink } from 'react-router-dom';

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
    marginLeft: 137,

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


const Login: FC<{}> = () => {
  const http = new DefaultApi();
  const classes = useStyles();
  const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);
  const [Status, setStatus] = React.useState(true);
  const [alert, setAlert] = React.useState(Boolean);

  const [Name, setName] = React.useState(String);
  const [Password, setPassword] = React.useState(String);
  const [users, setUsers] = React.useState<EntUser[]>([]);
  const arr = React.useState([]);
  const getUsers = async () => {
    const res = await http.listUser({ limit: 20, offset: 0 });
    setUsers(res);
  };
  //console.log(users);

  useEffect(() => {
    getUsers();
  }, []);

  // set data to object Patientright


  const Login = async () => {
    
    users.map((item) => {
      if(item.email === Name && item.password === Password){
      setCookie('Name', Name, { path: '/' })
      setCookie('Log', "true",{ path: '/' })
      setCookie('Status', item.edges?.userstatus?.userstatus , { path: '/' })

      setCookie('Fin', item.edges?.financier?.id , { path: '/' })
      setCookie('Med', item.edges?.medicalrecordstaff?.id , { path: '/' })
      setCookie('Nur', item.edges?.nurse?.id , { path: '/' })
      //setCookie('Doc', item.edges?.userPatientrights?.id , { path: '/' })
      //setCookie('Reg', item.edges?.user2registrar?.id , { path: '/' })
      window.location.reload(false)                
    }
    else{
      setStatus(true);
        setAlert(true);
        setTimeout(() => {
          setStatus(false);
        }, 1000);            
    }
  });
  }


  const NameChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setName(event.target.value as string);
    console.log(Name);
  };
  const PasswordChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPassword(event.target.value as string);
  };


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
              Login
            </Typography>
          </FormControl>
        </form>
        <form className={classes.form} noValidate>
          <FormControl variant="outlined" className={classes.formControl}>
            <TextField
              name="Name"
              label="Name"
              variant="outlined"
              type="string"
              size="medium"

              value={Name}
              onChange={NameChange}
            />&emsp;

            <TextField
            
              name="Password"
              label="Password"
              variant="outlined"
              type="password"
              size="medium"

              value={Password}
              onChange={PasswordChange}
            />

 
          </FormControl>
        </form>
        <form className={classes.form} noValidate>
        {Status ? (
                  <div>
                    {alert ? (
                        <Alert severity="error">
                            Email or Password not correct ,Try Again
                        </Alert>
                        ) : null}
                </div>
                 ) : null}
            <FormControl variant="outlined" className={classes.formControl}>
              <Button
                onClick={() => {
                  
                  Login();
                }}
                variant="contained"
                color="primary"
              >
                 Login
             </Button>
            </FormControl>
           &emsp;
             

        </form>
      </div>
    </Container>
  );
};
export default Login;