import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

import { ControllersDoctor } from '../../api/models/ControllersDoctor';
import { EntUser } from '../../api/models/EntUser';

import { EntDoctorinfo } from '../../api/models/EntDoctorinfo';
import { Alert } from '@material-ui/lab';
import { Cookies } from 'react-cookie/cjs';//cookie
import { Link as RouterLink } from 'react-router-dom';
// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');

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
  const http = new DefaultApi();
  const cookies = new Cookies();
  const [Doctor, setDoctor] = React.useState<Partial<ControllersDoctor>>({});
  const [User, setUser] = React.useState<EntUser[]>([]);
  const [Doctorinfo, setDoctorinfo] = React.useState<EntDoctorinfo[]>([]);
  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();
  const Img = cookies.get('Img');


  const getUser = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUser(res);
  };

  const getDoctorinfo = async () => {
    const res = await http.listDoctorinfo({ limit: 10, offset: 0 });
    setDoctorinfo(res);
  };

  const getImg = async () => {
    const res = await http.getUser({ id: Number(Img) });
    setUsers(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getUser();
    getDoctorinfo();
    getImg();
  }, []);

  // set data to object Patientright
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewPatientright;
    const { value } = event.target;
    setDoctor({ ...Doctor, [name]: value });
  };

  const CreatePatientright = async () => {

    if ((Doctor.doctorinfo != null) && (Doctor.user != null)
    ) {
      const res: any = await http.createDoctor({
        doctor: Doctor
      });
      console.log(Doctor);

      if (res.id != '') {
        setStatus(true);
        setAlert(true);
        setTimeout(() => {
          setStatus(false);
        }, 5000);   
      }
      window.location.reload(false)
    }
    else {
      setStatus(true);
      setAlert(false);
      setTimeout(() => {
        setStatus(false);
      }, 5000);
    }
  };


  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ลงทะหมอ`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
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
                    บันทึกไม่สำเร็จ
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>user</InputLabel>
              <Select
                name="user"
                value={Doctor.user}
                onChange={handleChange}
              >
                {User.map(item => {
                  if (item?.edges?.edgesOfDoctor?.id  >= 0) {
                    return null;
                  }
                  else {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.email} {item.edges?.edgesOfDoctor?.id}
                      </MenuItem>
                    );
                  }
                }
                )}
              </Select>
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>ข้อมูลหมอ</InputLabel>
              <Select
                name="doctorinfo"
                value={Doctor.doctorinfo}
                onChange={handleChange}
              >
                {Doctorinfo.map(item => {
                  if (item?.edges?.edgesOfDoctor?.id >= 0) {
                    return null;
                  }
                  else {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.doctorname} {item.doctorsurname} {item.edges?.edgesOfDoctor?.id}
                      </MenuItem>
                    );
                  }
                })}
              </Select>
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreatePatientright();
                }}
                variant="contained"
                color="primary"
              >
                Submit
             </Button>
             {/*
              <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                Back
                </Button>
                
              <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/Table_doctor"
                variant="contained"
                color="secondary"
              >
                SHOW
                </Button>
                */}
            </div>
          </form>
        </div>

      </Content>
    </Page>
  );
};
export default NewPatientright;