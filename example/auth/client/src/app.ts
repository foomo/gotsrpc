// Authentication state
let authToken: string | null = null;
let currentUser: any = null;

// DOM elements
const authStatus = document.getElementById('authStatus') as HTMLElement;
const loginForm = document.getElementById('loginForm') as HTMLElement;
const logoutForm = document.getElementById('logoutForm') as HTMLElement;
const authResult = document.getElementById('authResult') as HTMLElement;
const helloResult = document.getElementById('helloResult') as HTMLElement;
const userInfoResult = document.getElementById('userInfoResult') as HTMLElement;
const sayHelloBtn = document.getElementById('sayHelloBtn') as HTMLButtonElement;
const getUserInfoBtn = document.getElementById('getUserInfoBtn') as HTMLButtonElement;
const loginBtn = document.getElementById('loginBtn') as HTMLButtonElement;
const logoutBtn = document.getElementById('logoutBtn') as HTMLButtonElement;

// Import generated gotsrpc clients and types
import { AuthServiceClient, HelloServiceClient } from './client_gen.js';
import * as types from './vo_gen.js';

// Transport function for gotsrpc
const transport = <T>(endpoint: string) => async (method: string, data: any[] = []): Promise<T> => {
    const url = `http://localhost:8080${endpoint}/${encodeURIComponent(method)}`;
    
    const options: RequestInit = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    };

    // Add authorization header if we have a token
    if (authToken) {
        options.headers = {
            ...options.headers,
            'Authorization': `Bearer ${authToken}`,
        } as Record<string, string>;
    }

    if (data.length > 0) {
        options.body = JSON.stringify(data);
    }

    try {
        console.log('Making gotsrpc call to:', url);
        console.log('Method:', method, 'Data:', data);
        
        const response = await fetch(url, options);
        const responseText = await response.text();
        
        console.log('Response status:', response.status);
        console.log('Response text:', responseText);
        
        if (!response.ok) {
            throw new Error(`HTTP ${response.status}: ${responseText}`);
        }

        return responseText ? JSON.parse(responseText) : {};
    } catch (error) {
        console.error('gotsrpc call failed:', error);
        throw error instanceof Error ? error : new Error(String(error));
    }
};

// Create gotsrpc clients
const authClient = new AuthServiceClient(transport('/auth'));
const helloClient = new HelloServiceClient(transport('/hello'));

// Authentication functions
async function login() {
    const username = (document.getElementById('username') as HTMLInputElement).value;
    const password = (document.getElementById('password') as HTMLInputElement).value;

    if (!username || !password) {
        showResult(authResult, 'Please enter both username and password', 'error');
        return;
    }

    try {
        const loginReq: types.LoginRequest = { username, password };
        const response = await authClient.login(loginReq);
        
        if (response.ret_1) {
            throw new Error(response.ret_1.toString());
        }
        
        authToken = response.ret.token;
        currentUser = response.ret.user;
        
        updateAuthUI();
        showResult(authResult, `Login successful!\nUser: ${currentUser.username}\nToken: ${authToken?.substring(0, 20)}...`, 'success');
    } catch (error) {
        showResult(authResult, `Login failed: ${error instanceof Error ? error.message : String(error)}`, 'error');
    }
}

async function logout() {
    if (!authToken) return;

    try {
        const error = await authClient.logout(authToken);
        if (error) {
            throw new Error(error.toString());
        }
        
        authToken = null;
        currentUser = null;
        
        updateAuthUI();
        showResult(authResult, 'Logout successful!', 'success');
    } catch (error) {
        showResult(authResult, `Logout failed: ${error instanceof Error ? error.message : String(error)}`, 'error');
    }
}

// Hello service functions
async function sayHello() {
    const message = (document.getElementById('message') as HTMLTextAreaElement).value;

    if (!message.trim()) {
        showResult(helloResult, 'Please enter a message', 'error');
        return;
    }

    try {
        // First call Context to authenticate
        await helloClient.context();
        
        const helloReq: types.HelloRequest = { message };
        const response = await helloClient.sayHello(helloReq);
        
        if (response.ret_1) {
            throw new Error(response.ret_1.toString());
        }
        
        showResult(helloResult, `Response: ${JSON.stringify(response.ret, null, 2)}`, 'success');
    } catch (error) {
        showResult(helloResult, `Say hello failed: ${error instanceof Error ? error.message : String(error)}`, 'error');
    }
}

async function getUserInfo() {
    try {
        // First call Context to authenticate
        await helloClient.context();
        
        const response = await helloClient.getUserInfo();
        
        if (response.ret_1) {
            throw new Error(response.ret_1.toString());
        }
        
        showResult(userInfoResult, `User Info: ${JSON.stringify(response.ret, null, 2)}`, 'success');
    } catch (error) {
        showResult(userInfoResult, `Get user info failed: ${error instanceof Error ? error.message : String(error)}`, 'error');
    }
}

// UI helper functions
function updateAuthUI() {
    if (authToken && currentUser) {
        authStatus.textContent = `Authenticated as: ${currentUser.username}`;
        authStatus.className = 'auth-status authenticated';
        loginForm.classList.add('hidden');
        logoutForm.classList.remove('hidden');
        sayHelloBtn.disabled = false;
        getUserInfoBtn.disabled = false;
    } else {
        authStatus.textContent = 'Not authenticated';
        authStatus.className = 'auth-status not-authenticated';
        loginForm.classList.remove('hidden');
        logoutForm.classList.add('hidden');
        sayHelloBtn.disabled = false; // Allow say hello even when not logged in
        getUserInfoBtn.disabled = true; // Keep user info protected
    }
}

function showResult(element: HTMLElement, message: string, type: 'success' | 'error' | 'info') {
    element.textContent = message;
    element.className = `result ${type}`;
    element.classList.remove('hidden');
}

// Initialize UI
updateAuthUI();

// Add event listeners for buttons
if (loginBtn) loginBtn.addEventListener('click', login);
if (logoutBtn) logoutBtn.addEventListener('click', logout);
if (sayHelloBtn) sayHelloBtn.addEventListener('click', sayHello);
if (getUserInfoBtn) getUserInfoBtn.addEventListener('click', getUserInfo);

// Add some example credentials info
const infoDiv = document.createElement('div');
infoDiv.className = 'result info';
infoDiv.innerHTML = `
Available test accounts:
• Username: alice, Password: password123
• Username: bob, Password: secret456  
• Username: admin, Password: admin789

This client uses generated gotsrpc TypeScript clients!
`;
document.querySelector('.section')?.appendChild(infoDiv);
