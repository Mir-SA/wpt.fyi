/**
 * Copyright 2019 The WPT Dashboard Project. All rights reserved.
 * Use of this source code is governed by a BSD-style license that can be
 * found in the LICENSE file.
 */

const log = require('debug')('wpt.fyi');
const logDevData = require('debug')('wpt.fyi:dev_data');
const { spawn } = require('child_process');
const process = require('process');
const { DevAppserver } = require('./dev_appserver.js');

/**
 * @param {DevAppserver} server
 */
function populate(server) {
  return new Promise(resolve => {
    const args = [
        'run',
        '../util/populate_dev_data.go',
        `--project=${server.config.project}`,
        `--datastore_host=${server.datastoreUrl.host}`,
        `--local_host=${server.url.host}`,
        `--remote_runs=false`,
        `--static_runs=true`,
    ];
    log('Running go ' + args.join(' '));
    const child = spawn('go', args);
    process.on('exit', () => {
      log('killing dev_data subprocess...');
      child.kill();
    });
    child.stderr.on('data', buffer => { logDevData(buffer.toString()); });
    child.on('exit', exitCode => {
      log(`populate_dev_data.go exited with code ${exitCode}`);
      resolve();
      if (exitCode) {
        throw `dev_data child process exited with ${exitCode}`;
      }
    });
  });
}
exports.populate = populate;
