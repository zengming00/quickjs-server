import * as os from 'os';

// setInterval(()=>{ // 不支持
// 	console.log('helloworld for go & quickjs', Math.random(), new Date());
// }, 3000);

console.log('helloworld for go & quickjs', Math.random(), new Date());
foo.myPrint();

function sleep(ms) {
    return new Promise(function (resolve, reject) {
        console.log('in promise');
        os.setTimeout(function () {
            resolve({ a: 123, b: true, c: [3, 5] });
        }, ms);
    });
}
async function testAsync() {
    const r = await sleep(3000);
    console.log('r:', JSON.stringify(r));
    throw new Error('test err');
}
testAsync().catch(function (err) {
    console.log('err:', err);
})