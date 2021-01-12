import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

import {

  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,

  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command


import { Alert } from '@material-ui/lab';


import {  Cookies  } from 'react-cookie/cjs';//cookie

import { Link as RouterLink } from 'react-router-dom';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },


  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },

}));


const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'สิทธ์' };
  const http = new DefaultApi();

  const cookies = new Cookies();

  const Name = cookies.get('Name');
  const Password = cookies.get('Password');

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);


  
  // Lifecycle Hooks
  useEffect(() => {
   
  }, []);







  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ลงทะเบียน ${profile.givenName || 'to Backstage'}`}
        subtitle="Some quick intro and links."
      ><Timer /></Header>
      <Content>
        <ContentHeader title="ข้อมูล">

          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <TextField
            disabled
            id="outlined-disabled"
            label="Name"
            defaultValue={Name}
            variant="outlined"
          />

          <TextField
            disabled
            id="outlined-disabled"
            label="Password"
            defaultValue={Password}
            variant="outlined"
          />

        </div>

      </Content>
    </Page>
  );
};
export default NewPatientright;