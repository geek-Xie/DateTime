import axios from 'axios'
import StorageService from '../service/StorageService';
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 10000 * 5,
    header: { Authorization: `Bearer ${StorageService.get(StorageService.USER_TOKEN)}` }
});

export default service;