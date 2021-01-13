import React, { useState, useEffect } from 'react';
//import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Page,
  //BootErrorPageProps,
  Header,
  pageTheme,
  //ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { Cookies } from 'react-cookie/cjs';//cookie
import { Image6Base64Function } from '../../image/Image6';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { EntPrename } from '../../api/models/EntPrename';
import { EntEducationlevel } from '../../api/models/EntEducationlevel';
import { EntDepartment} from '../../api/models/EntDepartment';
import { EntOfficeroom } from '../../api/models/EntOfficeroom';
import { Grid,TextField, Avatar } from '@material-ui/core';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    formControl: {
        width: 250,
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

interface Doctorinfo_Type {
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  department?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  doctorname?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  doctorsurname?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  educationlevel?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  licensenumber?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  officeroom?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  prename?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  
  telephonenumber?: string;
}

export default function CreateDoctorinfo() {
  const classes = useStyles();
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [prenames, setPrenames] = useState<EntPrename[]>([]);
  const [educationlevels, setEducationlevels] = useState<EntEducationlevel[]>([]);
  const [departments, setDepartments] = useState<EntDepartment[]>([]);
  const [officerooms, setOfficerooms] = useState<EntOfficeroom[]>([]);

  const [Doctorinfo, setDoctorinfo] =            React.useState<Partial<Doctorinfo_Type>>({});

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getPrenames = async () => {
      const res = await api.listPrename({ limit: 5, offset: 0 });
      setLoading(false);
      setPrenames(res);
    };
    getPrenames();

    const getEducationlevels = async () => {
      const res = await api.listEducationlevel({ limit: 3, offset: 0 });
      setLoading(false);
      setEducationlevels(res);
      console.log(res);
    };
    getEducationlevels();

    const getDepartments = async () => {
      const res = await api.listDepartment({ limit: 3, offset: 0 });
      setLoading(false);
      setDepartments(res);
    };
    getDepartments();

    const getOfficerooms = async () => {
      const res = await api.listOfficeroom({ limit: 5, offset: 0 });
      setLoading(false);
      setOfficerooms(res);
    };
    getOfficerooms();

  }, [loading]);

const handleChange = (

  event: React.ChangeEvent<{ name?: string; value: unknown }>,
) => {
  const name = event.target.name as keyof typeof CreateDoctorinfo;
  const { value } = event.target;
  setDoctorinfo({ ...Doctorinfo, [name]: value });
};

  const Create_Doctorinfo = async () => {
   
    if ((Doctorinfo.department != null) && (Doctorinfo.doctorname != '')
      && (Doctorinfo.doctorsurname != '') && (Doctorinfo.educationlevel != null)
      && (Doctorinfo.licensenumber != '')&& (Doctorinfo.officeroom != null)
      && (Doctorinfo.prename != null)&& (Doctorinfo.telephonenumber != '')) {

    const res: any = await api.createDoctorinfo({ 
      doctorinfo:Doctorinfo
    
    
    });
    console.log(Doctorinfo);
    
    if (res.id != '') {
      setStatus(true);
      setAlert(true);
      setTimeout(() => {
        setStatus(false);
      }, 5000);
      }
      
    }
    else {
      setStatus(true);
      setAlert(false);
      setTimeout(() => {
        setStatus(false);
      }, 5000);
    }
  };

  const profile = { givenName: '' };
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`DOCTOR INFORMATION DEPARTMENT`}>
        <Avatar alt="Remy Sharp" src={Image6Base64Function} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <Grid container spacing = {5} >
          <Grid container item xs = {12} sm = {12}  >
              <Grid item xs = {12}>
                <Typography align ="center">
                    <Typography align = "center" variant = "h3">
                      <br/>=====  Create Doctorinformation  ===== <br/> <br/>
                    </Typography> 
                    <Typography variant="h6" gutterBottom  align="center">
                     เลือกคำนำหน้าชื่อ
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="prenames"
                    id="prenames"
                    name="prename"
                    value={Doctorinfo.prename}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {prenames.map((item: any) => (
                  <MenuItem value={item.id}>{item.prefix}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>

                
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="doctorname"
                                    label="ชื่อ"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.doctorname}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
                        
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="doctorsurname"
                                    label="นามสกุล"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.doctorsurname}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>

                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="telephonenumber"
                                    label="เบอร์ติดต่อ"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.telephonenumber}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
       
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="licensenumber"
                                    label="หมายเลขใบประกอบวิชาชีพ"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.licensenumber}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
                </Typography>

                <Typography variant="h6" gutterBottom  align="center">
                    เลือกระดับการศึกษา
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="educationlevels"
                    id="educationlevels"
                    name="educationlevel"
                    value={Doctorinfo.educationlevel}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {educationlevels.map((item: EntEducationlevel) => (
                  <MenuItem value={item.id}>{item.level}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>

                <Typography variant="h6" gutterBottom  align="center">
                    เลือกเเผนกของคุณ
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="departments"
                    id="departments"
                    name="department"
                    value={Doctorinfo.department}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {departments.map((item: any) => (
                    <MenuItem value={item.id}>{item.department}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>

                <Typography variant="h6" gutterBottom  align="center">
                    เลือกห้องพักของคุณ 
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="officerooms"
                    id="officerooms"
                    name="officeroom"
                    value={Doctorinfo.officeroom}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {officerooms.map((item: any) => (
                  <MenuItem value={item.id}>{item.roomnumber}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>
                <Grid container justify="center">
                    {status ? (
                    <div>
                        {alert ? (
                        <Alert severity="success">
                            success!
                        </Alert>
                        ) : (
                            <Alert severity="error" style={{ marginTop: 40 }}>
                            Failed to save!!! Please check again.
                            </Alert>
                        )}
                    </div>
                    ) : null}
                </Grid> <br/>

                <div className={classes.margin}>
                <Typography variant="h6" gutterBottom  align="center">
                <Button
                    onClick={() => {
                    Create_Doctorinfo();
                    }}
                    variant="contained"
                    color="primary"
                >
                    Submit
                </Button>
                </Typography> <br/>
                </div>
                </Grid>
            </Grid>
        </Grid>     
      </Content>
    </Page>
  );
};
