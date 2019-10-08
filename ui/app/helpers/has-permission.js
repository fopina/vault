import Helper from '@ember/component/helper';
import { inject as service } from '@ember/service';
import { observer } from '@ember/object';

export default Helper.extend({
  permissions: service(),
  /* eslint-disable-next-line ember/no-observers */
  onPermissionsChange: observer(
    'permissions.exactPaths',
    'permissions.globPaths',
    'permissions.canViewAll',
    function() {
      this.recompute();
    }
  ),

  compute([route], params) {
    let { routeParams, capabilities } = params;
    let permissions = this.permissions;
    if (route === 'access') {
      debugger;
    }
    return permissions.hasNavPermission(route, routeParams, capabilities);
  },
});
